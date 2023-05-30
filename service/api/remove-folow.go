package api

import (
	"net/http"

	"fmt"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from the follower list of another
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	oldFollower := ps.ByName("following_id")
	userPostId := ps.ByName("user_id")

	// Check if the id of the follower in the path is the same of bearer (no impersonation)
	valid := validateRequestingUser(oldFollower, requestingUserId)
	if valid != 0 {
		fmt.Println("ciao")
		w.WriteHeader(valid)
		return
	}

	if userPostId == requestingUserId {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserId}.ToDatabase(),
		UserId{User_id: userPostId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/rt.db.BanCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Remove the follower in the db via db function
	err = rt.db.UnfollowUser(
		UserId{User_id: oldFollower}.ToDatabase(),
		UserId{User_id: userPostId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-follow: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
