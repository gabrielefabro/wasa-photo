package api

import (
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a comment from a photo
func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserID := extractBearer(r.Header.Get("Authorization"))
	userID := ps.ByName("user_id")
	postIDStr := ps.ByName("post_id")
	commentIDStr := ps.ByName("comment_id")

	// Check if the user is not logged in
	if isNotLogged(requestingUserID) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if the requesting user is banned by the photo owner
	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserID}.ToDatabase(),
		UserId{User_id: userID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: BanCheck")
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Get the post ID from the path parameters
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to convert post_id to int64")
		return
	}

	// Get the comment ID from the path parameters
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to convert comment_id to int64")
		return
	}

	err = rt.db.UncommentPost(
		PostId{Post_id: postID}.ToDatabase(),
		UserId{User_id: requestingUserID}.ToDatabase(),
		CommentId{Comment_id: commentID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: UncommentPost")
		return
	}

	// Return a successful response with status code 204 (No Content)
	w.WriteHeader(http.StatusNoContent)
}
