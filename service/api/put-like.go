package api

import (
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that add a like of a user to a photo
func (rt *_router) putLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	postAuthor := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	pathLikeId := ps.ByName("like_id")

	// Check if the user is logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// User is trying to like his/her photo
	if postAuthor == requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BanCheck(
		User{User_id: requestingUserId}.ToDatabase(),
		User{User_id: postAuthor}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		// User was banned by owner, can't post the comment
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Follower id is not consistent with requesting user bearer token
	if pathLikeId != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert the path parameters photo_id from string to uint64
	post_id_64, err := strconv.ParseInt(ps.ByName("post_id"), 10, 64)
	post_id_u64 := uint64(post_id_64)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error converting path param photo_id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the like in the db via db function
	err = rt.db.LikePhoto(
		PostId{Post_id: post_id_u64}.ToDatabase(),
		User{User_id: pathLikeId}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
