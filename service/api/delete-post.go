package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that deletes a photo (this includes comments and likes)
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerAuth := extractBearer(r.Header.Get("Authorization"))
	postIdStr := ps.ByName("post_id")

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), bearerAuth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Convert the photo id from string to int64
	postInt, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-delete/ParseInt: error converting postId to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postUInt := uint64(postInt)

	// Call to the db function to remove the photo
	err = rt.db.DeletePhoto(
		User{User_id: bearerAuth}.ToDatabase(),
		PostId{Post_id: postUInt}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/RemovePhoto: error coming from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the folder of the file that has to be eliminated
	pathPhoto, err := getUserPhotoFolder(bearerAuth)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/getUserPhotoFolder: error with directories")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Remove the file from the user's photos folder
	err = os.Remove(filepath.Join(pathPhoto, postIdStr))
	if err != nil {
		// Error occurs if the file doesn't exist, but for idempotency an error won't be raised
		ctx.Logger.WithError(err).Error("photo-delete/os.Remove: photo to be removed is missing")
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
