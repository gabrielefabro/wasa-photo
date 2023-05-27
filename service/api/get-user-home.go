package api

import (
	"encoding/json"
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
		w.WriteHeader(valid)
		return
	}

	followers, err := rt.db.GetFollowings(User.ToDatabase(User{User_id: identifier}))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var posts []database.Post
	for _, follower := range followers {

		followerPost, err := rt.db.GetPosts(
			User.ToDatabase(User{}),
			follower)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		posts = append(posts, followerPost...)

	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(posts)
}
