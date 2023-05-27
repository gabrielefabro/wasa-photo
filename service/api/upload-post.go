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

// Function that manages the upload of a post
func (rt *_router) postPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	auth := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("user_id"), auth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	username, err := rt.db.GetUserName(auth)
	if err != nil {
		return
	}

	// Initialize photo struct
	post := Post{
		User_id:          auth,
		Username:         username,
		Publication_time: time.Now().UTC(),
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-upload: error reading body content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	err = checkFormatPhoto(r.Body, io.NopCloser(bytes.NewBuffer(data)), ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-upload: body contains file that is neither jpg or png")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: IMG_FORMAT_ERROR_MSG})
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	postIdInt, err := rt.db.UploadPost(post.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("function call")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	post_id := strconv.FormatInt(postIdInt, 10)

	// Create the user's folder locally to save his/her images
	PhotoPath, err := getUserPhotoFolder(auth)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-upload: error getting user's post folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create an empty file for storing the body content (image)
	out, err := os.Create(filepath.Join(PhotoPath, post_id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-upload: error creating local post file")
		//  = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
		return
	}

	// Copy body content to the previously created file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-upload: error copying body content into file photo")
		// controllaerrore
		// _ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
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
		Username:         post.Username,
		Publication_time: post.Publication_time,
		Post_id:          post.Post_id,
	})

}

// Function checks if the format of the post is png or jpeg. Returns the format extension and an error
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
func getUserPhotoFolder(user_id string) (UserPhotoFoldrPath string, err error) {

	// Path of the photo dir "./media/user_id/posts/"
	photoPath := filepath.Join(photoFolder, user_id, "posts")

	return photoPath, nil
}
