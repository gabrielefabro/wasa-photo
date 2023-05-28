package api

import (
	"net/http"
	"path/filepath"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

/*
getPosts is the handler for the GET /users/:profileUserID/posts endpoint
It return the posts of the user with the given profileUserID.
*/
func (rt *_router) getPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	requestedUser := ps.ByName("user_id")

	userBanned, err := rt.db.BanCheck(User{User_id: requestingUserId}.ToDatabase(),
		User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BanCheck/userBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	requestedProfileBanned, err := rt.db.BanCheck(User{User_id: requestedUser}.ToDatabase(),
		User{User_id: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BanCheck/requestedProfileBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if requestedProfileBanned {
		w.WriteHeader(http.StatusPartialContent)
		return
	}

	userExists, err := rt.db.CheckUser(User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.CheckUser: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !userExists {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Get the posts from the database
	dbPosts, err := rt.db.GetPosts(User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting posts")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	posts := make([]Post, len(dbPosts))

	for i, dbPosts := range dbPosts {
		var post Post
		err := post.FromDatabase(dbPosts)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error while converting the post")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		posts[i] = post
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
