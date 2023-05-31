package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from the follower list of another
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserID := extractBearer(r.Header.Get("Authorization"))
	followerID := ps.ByName("following_id")
	userID := ps.ByName("user_id")

	// Check if the follower ID in the path is the same as the requesting user (no impersonation)
	valid := validateRequestingUser(followerID, requestingUserID)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Users can't unfollow themselves
	if userID == requestingUserID {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Check if the requesting user is banned by the user being followed
	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserID}.ToDatabase(),
		UserId{User_id: userID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: BanCheck")
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Remove the follower from the database
	err = rt.db.UnfollowUser(
		UserId{User_id: followerID}.ToDatabase(),
		UserId{User_id: userID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: UnfollowUser")
		return
	}

	// Return a successful response with status code 204 (No Content)
	w.WriteHeader(http.StatusNoContent)
}
