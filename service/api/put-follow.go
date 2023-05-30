package api

import (
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to the followers list of another
func (rt *_router) putFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userToFollowId := ps.ByName("user_id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// users can't follow themselves
	if requestingUserId == userToFollowId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the id of the follower in the request is the same of the bearer and the path parameter
	if ps.ByName("follower_id") != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserId}.ToDatabase(),
		UserId{User_id: userToFollowId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/rt.db.BanCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Add the new follower in the db via db function
	err = rt.db.FollowUser(
		UserId{User_id: requestingUserId}.ToDatabase(),
		UserId{User_id: userToFollowId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
