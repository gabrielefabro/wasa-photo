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
	requestedUser := ps.ByName("user_id")

	var followers []database.User
	var following []database.User
	var posts []database.Post

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

	followers, err = rt.db.GetFollowers(User{User_id: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetMyFollowers: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	following, err = rt.db.GetFollowings(User{User_id: requestedUser}.ToDatabase())
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

	username, err := rt.db.GetUserName(requestedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetuserName: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Profile{
		User_id:   requestedUser,
		Username:  username,
		Follower:  followers,
		Following: following,
		Posts:     posts,
	})

}
