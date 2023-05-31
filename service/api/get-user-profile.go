package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"
	"git.gabrielefabro.it/service/database"

	"github.com/julienschmidt/httprouter"
)

// Function that retrieves all the necessary information of a profile
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	requestedUser := ps.ByName("user_id")

	var followers []database.UserId
	var following []database.UserId
	var posts []database.Post

	// Check if the requesting user is banned
	userBanned, err := rt.db.BanCheck(UserId{User_id: requestingUserId}.ToDatabase(), UserId{User_id: requestedUser}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: getUserProfile/db.BanCheck/userBanned")
		return
	}
	if userBanned {
		handleError(w, http.StatusForbidden, "")
		return
	}

	// Check if the requested user's profile is banned
	requestedProfileBanned, err := rt.db.BanCheck(UserId{User_id: requestedUser}.ToDatabase(), UserId{User_id: requestingUserId}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: getUserProfile/db.BanCheck/requestedProfileBanned")
		return
	}
	if requestedProfileBanned {
		handleError(w, http.StatusPartialContent, "")
		return
	}

	// Check if the requested user exists
	userExists, err := rt.db.CheckUser(UserId{User_id: requestedUser}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: getUserProfile/db.CheckUser")
		return
	}
	if !userExists {
		handleError(w, http.StatusNoContent, "")
		return
	}

	// Get the followers of the requested user
	followers, err = rt.db.GetFollowers(UserId{User_id: requestedUser}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: getUserProfile/db.GetFollowers")
		return
	}

	// Get the users that the requested user is following
	following, err = rt.db.GetFollowings(UserId{User_id: requestedUser}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: getUserProfile/db.GetFollowings")
		return
	}

	// Get the posts of the requested user
	posts, err = rt.db.GetPosts(UserId{User_id: requestingUserId}.ToDatabase(), UserId{User_id: requestedUser}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: getUserProfile/db.GetPosts")
		return
	}

	// Get the username of the requested user
	username, err := rt.db.GetUserName(UserId{User_id: requestedUser}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: getUserProfile/db.GetUserName")
		return
	}

	handleSuccess(w, http.StatusOK, Profile{
		User_id:   requestedUser,
		Username:  username,
		Follower:  followers,
		Following: following,
		Posts:     posts,
	})
}
