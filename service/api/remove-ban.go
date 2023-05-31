package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from another user's banned list
func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserID := extractBearer(r.Header.Get("Authorization"))
	pathUserID := ps.ByName("user_id")
	bannedUserID := ps.ByName("banned_id")

	valid := validateRequestingUser(pathUserID, requestingUserID)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Users cannot unban themselves
	if bannedUserID == requestingUserID {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err := rt.db.UnbanUser(
		UserId{User_id: pathUserID}.ToDatabase(),
		UserId{User_id: bannedUserID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute delete query: UnbanUser")
		return
	}

	// Return a successful response with status code 204 (No Content)
	w.WriteHeader(http.StatusNoContent)
}
