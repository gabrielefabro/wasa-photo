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
	postUserId := ps.ByName("user_id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserId}.ToDatabase(),
		UserId{User_id: postUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BanCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment/Decode: failed to decode request body json")
		return
	}

	// Check if the comment has a valid lenght (<=50)
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

	// Function call to db for comment creation
	commentId, err := rt.db.CommentPost(
		PostId{Post_id: post_id_64}.ToDatabase(),
		UserId{User_id: requestingUserId}.ToDatabase(),
		TextComment{TextComment: comment.Text}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/db.CommentPost: failed to execute query for insertion")
		return
	}

	w.WriteHeader(http.StatusCreated)

	// The response body will contain the unique id of the comment
	err = json.NewEncoder(w).Encode(CommentId{Comment_id: commentId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/Encode: failed convert post_id to int64")
		return
	}
}
