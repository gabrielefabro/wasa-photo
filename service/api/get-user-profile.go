package api

import (
	"encoding/json"
	"net/http"

	"git.gabrielefabro.it/service/api/reqcontext"
	"git.gabrielefabro.it/service/database"

	"github.com/julienschmidt/httprouter"
)

// Function that retrives all the necessary infos of a profile
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	requestedUser := ps.ByName("id")

	var followers []database.User
	var following []database.User
	var posts []database.Post

	// Check if the requesting user is banned by the requested profile owner
	userBanned, err := rt.db.BanCheck(User{User_id: requestingUserId}.ToDatabase(),
		User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BanCheck/userBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if the requested profile was banned by the requesting user. If it's true respond with partial content
	requestedProfileBanned, err := rt.db.BanCheck(User{User_id: requestedUser}.ToDatabase(),
		User{User_id: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BanCheck/requestedProfileBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if requestedProfileBanned {
		w.WriteHeader(http.StatusPartialContent)
		return
	}

	userExists, err := rt.db.CheckUser(User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.CheckUser: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !userExists {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	followers, err = rt.db.GetMyFollowers(User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetMyFollowers: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	following, err = rt.db.GetMyFollowings(User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetMyFollowings: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	posts, err = rt.db.GetPosts(User{User_id: requestingUserId}.ToDatabase(), User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetPosts: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	username, err := rt.db.GetUserName(User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetuserName: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Profile{
		User:      User{User_id: requestedUser, UserName: username},
		Follower:  followers,
		Following: following,
		Posts:     posts,
	})

}
