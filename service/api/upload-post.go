package api

import (
	"encoding/json"
	"io"
	"net/http"
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

	// Initialize photo struct
	photo := Photo{
		user_id: auth,
		pubblication_time:  time.Now().UTC(),
	}

	// Create a copy of the body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-upload: error reading body content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// After reading the body we won't be able to read it again. We'll reassign a "fresh" io.ReadCloser to the body
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// Check if the body content is either a png or a jpeg image
	err = checkFormatPhoto(r.Body, io.NopCloser(bytes.NewBuffer(data)), ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-upload: body contains file that is neither jpg or png")
		// controllaerrore
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: IMG_FORMAT_ERROR_MSG})
		return
	}

	// Body has been read in the previous function so it's necessary to reassign a io.ReadCloser to it
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// Generate a unique id for the photo
	postIdInt, err := rt.db.UploadPost(post.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error executing db function call")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	post_id := strconv.FormatInt(postIdInt, 10)

	// Create the user's folder locally to save his/her images
	PhotoPath, err := getUserPhotoFolder(auth)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error getting user's photo folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create an empty file for storing the body content (image)
	out, err := os.Create(filepath.Join(PhotoPath, post_id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error creating local photo file")
		//  = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
		return
	}

	// Copy body content to the previously created file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error copying body content into file photo")
		// controllaerrore
		// _ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
		return
	}

	// Close the created file
	out.Close()

	username, err := db.GetUserName(post.User_id)
		if err != nil {
			return nil, err
		}
		comment.Username = username

	w.WriteHeader(http.StatusCreated)
	// controllaerrore
	// _ = json.NewEncoder(w).Encode(PhotoId{IdPhoto: photoIdInt})
	_ = json.NewEncoder(w).Encode(Photo{
		Comments: nil,
		Likes:    nil,
		user_id:  post.user_id,
		Date:     post.pubblication_time,
		post_id:  post_id,
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
func getUserPhotoFolder(user_id string) (UserPhotoFoldrPath string, err error) {

	// Path of the photo dir "./media/user_id/photos/"
	photoPath := filepath.Join(photoFolder, user_id, "photos")

	return photoPath, nil
}