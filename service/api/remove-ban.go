package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from the banned list of another
func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerToken := extractBearer(r.Header.Get("Authorization"))
	pathId := ps.ByName("user_id")
	userToUnban := ps.ByName("banned_id")

	valid := validateRequestingUser(pathId, bearerToken)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Users can't ban themselfes
	if userToUnban == bearerToken {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Remove the follower in the db via db function
	err := rt.db.UnbanUser(
		UserId{User_id: pathId}.ToDatabase(),
		UserId{User_id: userToUnban}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-ban: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
