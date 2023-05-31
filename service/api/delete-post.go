package api

import (
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that deletes a post
func (rt *_router) deletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerAuth := extractBearer(r.Header.Get("Authorization"))
	postIdStr := ps.ByName("post_id")

	valid := validateRequestingUser(ps.ByName("user_id"), bearerAuth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	postInt, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error converting post ID to int")
		return
	}

	err = rt.db.DeletePost(UserId{User_id: bearerAuth}.ToDatabase(), PostId{Post_id: postInt}.ToDatabase())
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Database error while deleting the post")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
