package api

import (
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that delete a post
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerAuth := extractBearer(r.Header.Get("Authorization"))
	postIdStr := ps.ByName("post_id")
	user_id := ps.ByName("user_id")

	// Check the user
	if errCode := validateRequestingUser(user_id, bearerAuth); errCode != 0 {
		w.WriteHeader(errCode)
		return
	}

	post_id, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		handleError(ctx, err, "post-delete/ParseUint: errore nella conversione di post_id in uint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := rt.db.DeletePost(User{User_id: bearerAuth}.ToDatabase(), PostId{Post_id: post_id}.ToDatabase()); err != nil {
		handleError(ctx, err, "post-delete/DeletePost: errore proveniente dal database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleError(ctx reqcontext.RequestContext, err error, message string) {
	ctx.Logger.WithError(err).Error(message)
}
