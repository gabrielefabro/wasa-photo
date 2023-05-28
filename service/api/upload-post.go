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
	user_id := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("user_id"), user_id)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	username, err := rt.db.GetUserName(user_id)
	if err != nil {
		return
	}

	// Parse the multipart form
	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parse file")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fileType := http.DetectContentType(data)
	if fileType != "image/jpeg" {
		http.Error(w, "Bad Request wrong file type", http.StatusBadRequest)
		return
	}

	defer func() { err = file.Close() }()

	// Get the username from the database
	username, err = rt.db.GetUserName(user_id)
	if err != nil {
		ctx.Logger.WithError(err).Error("error getting user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create a new post
	var post = Post{
		User_id:          user_id,
		Username:         username,
		Publication_time: time.Now().UTC(),
	}

	dbPost := post.ToDatabase()

	post_id, err := rt.db.UploadPost(dbPost, data)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating post")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	post.Post_id = uint64(post_id)
	// Parse the new post from the database package to the Post struct in the api package
	err = post.FromDatabase(dbPost)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parsing photo")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the new post
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the post")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
