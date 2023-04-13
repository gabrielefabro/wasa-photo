package api

import (
	"encoding/json"
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"
	"git.gabrielefabro.it/service/database"

	"github.com/julienschmidt/httprouter"
)

// This function retrieves all the photos of the people that the user is following
func (rt *_router) getHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	identifier := extractBearer(r.Header.Get("Authorization"))

	// A user can only see his/her home
	valid := validateRequestingUser(ps.ByName("id"), identifier)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	followers, err := rt.db.GetMyFollowings(User.ToDatabase(User{User_id: identifier}))
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

		for _, post := range followerPost {
			posts = append(posts, post)
		}

	}

	w.WriteHeader(http.StatusOK)

	// Send the output to the user. Instead of giving null for no matches return and empty slice of photos.
	_ = json.NewEncoder(w).Encode(posts)
}
