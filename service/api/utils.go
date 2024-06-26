package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Function that verifies if the identifier of a user has the right lenght
func validIdentifier(identifier string) bool {
	var trimmedId = strings.TrimSpace(identifier)
	return len(identifier) >= 3 && len(identifier) <= 16 && trimmedId != "" && !strings.ContainsAny(trimmedId, "?_")
}

// Function that extracts the bearer token from the Authorization header
func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}

// Utility function to handle errors
func handleError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	if message != "" {
		_ = json.NewEncoder(w).Encode(struct{ Message string }{Message: message})
	}
}

// Utility function to handle successful responses
func handleSuccess(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// Function that checks if the requesting user has a valid token for the specified endpoint. Returns 0 if it's valid, the error (as a int, representing the http status) otherwise
func validateRequestingUser(identifier string, bearerToken string) int {

	// If the requesting user has an invalid token then respond with a fobidden status
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}
	if identifier != bearerToken {
		return http.StatusUnauthorized
	}
	return 0
}

// Function that checks if a user is logged
func isNotLogged(auth string) bool {

	return auth == ""
}
