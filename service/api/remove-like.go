package api

import (
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a like from a photo
func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	postAuthor := ps.ByName("user_id")
	requestingUserID := extractBearer(r.Header.Get("Authorization"))

	// Check if the user is logged
	if isNotLogged(requestingUserID) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if the requesting user is banned by the photo owner
	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserID}.ToDatabase(),
		UserId{User_id: postAuthor}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: BanCheck")
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	postID, err := strconv.ParseInt(ps.ByName("post_id"), 10, 64)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to convert post_id to int64")
		return
	}

	// Remove the like from the database
	err = rt.db.UnlikePost(
		PostId{Post_id: postID}.ToDatabase(),
		UserId{User_id: requestingUserID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: UnlikePost")
		return
	}

	// Return a successful response with status code 204 (No Content)
	w.WriteHeader(http.StatusNoContent)
}
