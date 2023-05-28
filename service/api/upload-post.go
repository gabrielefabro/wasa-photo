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
	user_id := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("user_id"), user_id)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	username, err := rt.db.GetUserName(auth)
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

	// Get the user from the database
	dbuser, err := rt.db.GetUserByID(user_id)
	if err != nil {
		ctx.Logger.WithError(err).Error("error getting user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the user from the database to the User struct in the api package
	var user User
	err = user.FromDatabase(dbuser)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parsing user")
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
	

	dbNewPost, err := rt.db.CreatePost(dbPost, data)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating post")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the new post from the database package to the Post struct in the api package
	err = newPost.FromDatabase(dbNewPost)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parsing photo")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the new post
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newPost); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the post")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}