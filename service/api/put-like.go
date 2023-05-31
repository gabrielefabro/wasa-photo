package api

import (
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user's like to a photo
func (rt *_router) putLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	postID, err := strconv.ParseInt(ps.ByName("post_id"), 10, 64)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to convert post_id to int64")
		return
	}

	requestingUserID := extractBearer(r.Header.Get("Authorization"))

	if isNotLogged(requestingUserID) {
		handleError(w, http.StatusForbidden, "User is not logged in")
		return
	}

	postAuthor := ps.ByName("user_id")

	banned, err := rt.db.BanCheck(
		UserId{User_id: requestingUserID}.ToDatabase(),
		UserId{User_id: postAuthor}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to check if user is banned")
		return
	}
	if banned {
		handleError(w, http.StatusForbidden, "User is banned")
		return
	}

	err = rt.db.LikePost(
		PostId{Post_id: postID}.ToDatabase(),
		UserId{User_id: requestingUserID}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to execute query: LikePost")
		return
	}

	// Return a successful response with status code 204 (No Content)
	w.WriteHeader(http.StatusNoContent)
}
