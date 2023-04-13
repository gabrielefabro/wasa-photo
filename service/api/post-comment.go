package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a comment to a photo and sends a response containing the unique id of the created comment
func (rt *_router) postComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	postUserId := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BanCheck(
		User{User_id: requestingUserId}.ToDatabase(),
		User{User_id: postUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BanCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		// User was banned by owner, can't post the comment
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Copy body content (comment sent by user) into comment (struct)
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment/Decode: failed to decode request body json")
		return
	}

	// Check if the comment has a valid lenght (<=30)
	if len(comment.Text) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: comment longer than 50 characters")
		return
	}

	// Convert the photo identifier from string to int64
	post_id_64, err := strconv.ParseInt(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment/ParseInt: failed convert photo_id to int64")
		return
	}

	// cast to make a variable int64 to uint64
	post_id_u64 := uint64(post_id_64)

	// Function call to db for comment creation
	commentId, err := rt.db.CommentPhoto(
		PostId{Post_id: post_id_u64}.ToDatabase(),
		User{User_id: requestingUserId}.ToDatabase(),
		TextComment{TextComment: comment.Text}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/db.CommentPhoto: failed to execute query for insertion")
		return
	}

	w.WriteHeader(http.StatusCreated)

	// cast to make a variable int64 to uint64
	commentIdU := uint64(commentId)

	// The response body will contain the unique id of the comment
	err = json.NewEncoder(w).Encode(CommentId{Comment_id: commentIdU})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/Encode: failed convert post_id to int64")
		return
	}
}
