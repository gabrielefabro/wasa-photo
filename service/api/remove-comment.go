package api

import (
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a comment from a photo
func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// Check if the user isn't logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserId}.ToDatabase(),
		UserId{User_id: ps.ByName("user_id")}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BanCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	post_id_64, err := strconv.ParseInt(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: failed convert photo_id to int64")
		return
	}

	// Convert the comment identifier from string to int64
	comment_id_64, err := strconv.ParseInt(ps.ByName("comment_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: failed convert comment_id to int64")
		return
	}

	// Function call to db for comment removal (only authors can remove their comments)
	err = rt.db.UncommentPost(
		PostId{Post_id: post_id_64}.ToDatabase(),
		UserId{User_id: requestingUserId}.ToDatabase(),
		CommentId{Comment_id: comment_id_64}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("delete-comment: failed to execute query for insertion")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
