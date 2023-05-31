package api

import (
	"encoding/json"
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that updates a user's nickname
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the user ID from the path
	pathId := ps.ByName("user_id")

	// Check the user's identity for the operation
	if valid := validateRequestingUser(pathId, extractBearer(r.Header.Get("Authorization"))); valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Decode the new username from the request body
	var username Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	// Modify the username using the db function
	err = rt.db.ChangeUserName(
		UserId{User_id: pathId}.ToDatabase(),
		username.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to update username")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
