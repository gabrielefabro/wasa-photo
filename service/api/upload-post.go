package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that manages the upload of a photo
func (rt *_router) postPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	auth := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), auth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}
	// Convert auth to uint64 from string
	authInt, err := strconv.ParseInt(auth, 10, 64)
	authUInt := uint64(authInt)

	// Initialize photo struct
	post := Post{
		User_id:          authUInt,
		Publication_time: time.Now().UTC(),
	}

	// Create a copy of the body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("uplad-post: error reading body content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// After reading the body we won't be able to read it again. We'll reassign a "fresh" io.ReadCloser to the body
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// Check if the body content is either a png or a jpeg image
	err = checkFormatPhoto(r.Body, io.NopCloser(bytes.NewBuffer(data)), ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("upload-post: body contains file that is neither jpg or png")
		// controllaerrore
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: IMG_FORMAT_ERROR_MSG})
		return
	}

	// Body has been read in the previous function so it's necessary to reassign a io.ReadCloser to it
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// Generate a unique id for the photo
	postIdInt, err := rt.db.UploadPhoto(post.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("upload-post: error executing db function call")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postId := strconv.FormatInt(postIdInt, 10)

	// Create the user's folder locally to save his/her images
	PostPath, err := getUserPostFolder(auth)
	if err != nil {
		ctx.Logger.WithError(err).Error("upload-post: error getting user's photo folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create an empty file for storing the body content (image)
	out, err := os.Create(filepath.Join(PostPath, postId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("upload-post: error creating local photo file")
		return
	}

	// Copy body content to the previously created file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("upload-post: error copying body content into file photo")
		return
	}

	// Close the created file
	out.Close()

	w.WriteHeader(http.StatusCreated)
	// controllaerrore
	// _ = json.NewEncoder(w).Encode(PhotoId{IdPhoto: photoIdInt})
	_ = json.NewEncoder(w).Encode(Post{
		Comment:          nil,
		Like:             nil,
		User_id:          post.User_id,
		Publication_time: post.Publication_time,
		Post_id:          uint64(postIdInt),
	})

}

// Function checks if the format of the photo is png or jpeg. Returns the format extension and an error
func checkFormatPhoto(body io.ReadCloser, newReader io.ReadCloser, ctx reqcontext.RequestContext) error {

	_, errJpg := jpeg.Decode(body)
	if errJpg != nil {

		body = newReader
		_, errPng := png.Decode(body)
		if errPng != nil {
			return errors.New(IMG_FORMAT_ERROR_MSG)
		}
		return nil
	}
	return nil
}

// Function that returns the path of the photo folder for a certain user
func getUserPostFolder(user_id string) (UserPhotoFoldrPath string, err error) {

	// Path of the photo dir "./media/user_id/photos/"
	photoPath := filepath.Join(postFolder, user_id, "photos")

	return photoPath, nil
}
