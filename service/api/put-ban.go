package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to the banned list of another user
func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	pathID := ps.ByName("user_id")
	pathBannedID := ps.ByName("banned_id")
	requestingUserID := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation
	if valid := validateRequestingUser(pathID, requestingUserID); valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Check if the user is trying to ban themselves
	if requestingUserID == pathBannedID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add the new banned user in the database
	if err := rt.db.BanUser(UserId{User_id: pathID}.ToDatabase(), UserId{User_id: pathBannedID}.ToDatabase()); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute insert query: put-ban/db.BanUser")
		return
	}

	// Remove the follow relationship (if it exists) between the users
	if err := rt.db.UnfollowUser(UserId{User_id: requestingUserID}.ToDatabase(), UserId{User_id: pathBannedID}.ToDatabase()); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute insert query: put-ban/db.UnfollowUser1")
		return
	}

	// Remove the follow relationship (if it exists) between the users in the reverse direction
	if err := rt.db.UnfollowUser(UserId{User_id: pathBannedID}.ToDatabase(), UserId{User_id: requestingUserID}.ToDatabase()); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute insert query: put-ban/db.UnfollowUser2")
		return
	}

	handleSuccess(w, http.StatusNoContent, nil)
}
