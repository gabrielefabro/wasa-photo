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
	postUserID := ps.ByName("user_id")
	requestingUserID := extractBearer(r.Header.Get("Authorization"))

	if isNotLogged(requestingUserID) {
		handleError(w, http.StatusForbidden, "")
		return
	}

	banned, err := rt.db.BanCheck(UserId{User_id: requestingUserID}.ToDatabase(), UserId{User_id: postUserID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing query: post-comment/db.BanCheck")
		return
	}
	if banned {
		handleError(w, http.StatusForbidden, "")
		return
	}

	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to decode request body JSON: post-comment/Decode")
		return
	}

	// Check if the comment has a valid length (<=50)
	if len(comment.Text) > 50 {
		handleError(w, http.StatusBadRequest, "Comment longer than 50 characters: post-comment")
		return
	}

	// Convert the post identifier from string to int64
	postID, err := strconv.ParseInt(ps.ByName("post_id"), 10, 64)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to convert post_id to int64: post-comment/ParseInt")
		return
	}

	// Call the database function to create the comment
	commentID, err := rt.db.CommentPost(PostId{Post_id: postID}.ToDatabase(), UserId{User_id: requestingUserID}.ToDatabase(), TextComment{TextComment: comment.Text}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query for comment insertion: post-comment/db.CommentPost")
		return
	}

	handleSuccess(w, http.StatusCreated, CommentId{Comment_id: commentID})
}
