package api

import (
	"encoding/json"
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that updates a user's nickname
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	pathId := ps.ByName("user_id")

	// Check the user's identity for the operation
	valid := validateRequestingUser(pathId, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Get the new username from the request body
	var username Username
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		ctx.Logger.WithError(err).Error("update-username: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Modify the username with the db function
	err = rt.db.ModifyUserName(
		UserId{User_id: pathId}.ToDatabase(),
		username.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("update-username/ModifyUsername: error executing update query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
