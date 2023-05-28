package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sessionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !validIdentifier(user.User_id) {
		// Here we checked the user identifier and we discovered that it's not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user_id = user.User_id
	// Create the user in the database.
	err = rt.db.CreateUser(user_id)
	if err != nil {
		// In this case, there's a sql error since the resource already exists and can't be inserted again.
		// The identifier is returned as expected.
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
		}
		return
	}

	// Create user's directories locally
	err = createUserFolder(user.User_id, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("session: can't create user's post folder")
		return
	}

	// Send the output to the user
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}

}

// Function that creates a new subdir for the specified user
func createUserFolder(identifier string, ctx reqcontext.RequestContext) error {

	// Create the path media/useridentifier/ inside the project dir
	path := filepath.Join(photoFolder, identifier)

	// To the previously created path add the "posts" subdir
	err := os.MkdirAll(filepath.Join(path, "posts"), os.ModePerm)
	if err != nil {
		ctx.Logger.WithError(err).Error("session/createUserFolder:: error creating directories for user")
		return err
	}
	return nil
}
