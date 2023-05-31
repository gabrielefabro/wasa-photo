package api

import (
	"encoding/json"
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that retrieves all the users matching the query parameter and sends the response containing all the matches
func (rt *_router) getUsersQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")

	// Get the user identifier (from Bearer)
	identifier := extractBearer(r.Header.Get("Authorization"))

	// If the user is not logged in, respond with a 403 HTTP status
	if identifier == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Extract the query parameter from the URL
	queryParam := r.URL.Query().Get("user_id")

	// Search the user in the database using the query parameter as a filter
	users, err := rt.db.SearchUser(UserId{User_id: identifier}.ToDatabase(), UserId{User_id: queryParam}.ToDatabase())
	if err != nil {
		// If there's an error from the database, return an empty JSON response
		handleError(w, http.StatusInternalServerError, "Database error")
		return
	}

	w.WriteHeader(http.StatusOK)

	// Send the output to the user. Instead of giving null for no matches return and empty slice of Users
	if len(users) == 0 {
		_ = json.NewEncoder(w).Encode([]UserId{})
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}
