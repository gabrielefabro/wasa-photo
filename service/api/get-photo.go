package api

import (
	"net/http"
	"path/filepath"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that serves the requested photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	http.ServeFile(w, r,
		filepath.Join(photoFolder, ps.ByName("user_id"), "posts", ps.ByName("post_id")))

}
