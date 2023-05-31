package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"
	"git.gabrielefabro.it/service/database"

	"github.com/julienschmidt/httprouter"
)

// This function retrieves all the posts of the people that the user is following
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	identifier := extractBearer(r.Header.Get("Authorization"))

	valid := validateRequestingUser(ps.ByName("user_id"), identifier)
	if valid != 0 {
		handleError(w, valid, "")
		return
	}

	followers, err := rt.db.GetFollowings(UserId{User_id: identifier}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "")
		return
	}

	var posts []database.Post
	for _, follower := range followers {
		followerPost, err := rt.db.GetPosts(UserId{User_id: identifier}.ToDatabase(), UserId{User_id: follower.User_id}.ToDatabase())
		if err != nil {
			handleError(w, http.StatusInternalServerError, "")
			return
		}
		posts = append(posts, followerPost...)
	}

	handleSuccess(w, http.StatusOK, posts)
}
