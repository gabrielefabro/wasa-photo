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
		filepath.Join(postFolder, ps.ByName("id"), "posts", ps.ByName("posts_id")))

}
