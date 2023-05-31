package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to the followers list of another user
func (rt *_router) putFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userToFollowID := ps.ByName("user_id")
	requestingUserID := extractBearer(r.Header.Get("Authorization"))

	// Users can't follow themselves
	if requestingUserID == userToFollowID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the follower ID in the request matches the bearer and path parameter
	if ps.ByName("follower_id") != requestingUserID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the requesting user has been banned by the user they want to follow
	if banned, err := rt.db.BanCheck(UserId{User_id: requestingUserID}.ToDatabase(), UserId{User_id: userToFollowID}.ToDatabase()); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: put-follow/rt.db.BanCheck")
		return
	} else if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Add the new follower in the database
	if err := rt.db.FollowUser(UserId{User_id: requestingUserID}.ToDatabase(), UserId{User_id: userToFollowID}.ToDatabase()); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute insert query: put-follow")
		return
	}

	handleSuccess(w, http.StatusNoContent, nil)
}
