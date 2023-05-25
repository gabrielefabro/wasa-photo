package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"
	"git.gabrielefabro.it/service/database"

	"github.com/julienschmidt/httprouter"
)

// Function that serves the requested photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerAuth := extractBearer(r.Header.Get("Authorization"))
	postIdStr := ps.ByName("post_id")
	user_id := ps.ByName("user_id")

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

	post, err := rt.db.GetPhoto(database.User{User_id: user_id}, database.PostId{Post_id: post_id})
	if err != nil {
		handleError(ctx, err, "photo-get/GetPhoto: errore proveniente dal database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(post)
}
