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
		ctx.Logger.WithError(err).Error("post-delete/ParseInt: error converting photoId to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Call to the db function to remove the post
	err = rt.db.DeletePost(
		User{User_id: bearerAuth}.ToDatabase(),
		PostId{Post_id: uint64(postInt)}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-delete/RemovePhoto: error coming from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
