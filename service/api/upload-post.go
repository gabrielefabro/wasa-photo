package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that manages the upload of a photo
func (rt *_router) postPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("user_id"), auth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		ctx.Logger.WithError(err).Error("Bad request, out of memory")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		ctx.Logger.WithError(err).Error("Bad request, can't load the image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		ctx.Logger.WithError(err).Error("Bad request, can't read the image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileType := http.DetectContentType(data)
	if fileType != "image/jpeg" {
		ctx.Logger.WithError(err).Error("Bad request, image format should be jpeg")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() { err = file.Close() }()

	// Get username for the post
	username, err := rt.db.GetUserName(auth)
	if err != nil {
		ctx.Logger.WithError(err).Error("Database error can't get the username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Create the post
	var post = Post{
		User_id:          auth,
		Username:         username,
		Publication_time: time.Now().UTC(),
	}

	// Generate a unique id for the photo
	postIdInt, err := rt.db.UploadPost(Post.ToDatabase(post))
	if err != nil {
		ctx.Logger.WithError(err).Error("upload-post: error executing db function call")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	postId := uint64(postIdInt)

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(Post{
		Comment:          nil,
		Like:             nil,
		User_id:          post.User_id,
		Publication_time: post.Publication_time,
		Post_id:          postId,
		Username:         username,
	})

}
