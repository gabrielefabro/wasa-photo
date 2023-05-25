package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login
	rt.router.POST("/session", rt.wrap(rt.sessionHandler))

	// Change UserName
	rt.router.PUT("/users/:user_id", rt.wrap(rt.setMyUserName))

	// Return Profile
	rt.router.GET("/users/:user_id", rt.wrap(rt.getUserProfile))

	// Stream
	rt.router.GET("/users/:user_id/home", rt.wrap(rt.getMyStream))

	// Follow User
	rt.router.PUT("/users/:user_id/followings/:following_id", rt.wrap(rt.putFollow))

	// Unfollow User
	rt.router.DELETE("/users/:user_id/followings/:following_id", rt.wrap(rt.deleteFollow))

	// Ban User
	rt.router.PUT("/users/:user_id/banned_users/:banned_id", rt.wrap(rt.putBan))

	// Unban User
	rt.router.DELETE("/users/:user_id/banned_users/:banned_id", rt.wrap(rt.deleteBan))

	// Like Post
	rt.router.PUT("/users/:user_id/posts/:post_id/likes/:like_id", rt.wrap(rt.putLike))

	// Unlike Post
	rt.router.DELETE("/users/:user_id/posts/:post_id/likes/:like_id", rt.wrap(rt.deleteLike))

	// Return Post
	rt.router.GET("/users/:user_id/posts/:post_id", rt.wrap(rt.getPhoto))

	// Delete Post
	rt.router.DELETE("/users/:user_id/posts/:post_id", rt.wrap(rt.deletePhoto))

	// Post
	rt.router.POST("/users/:user_id/posts", rt.wrap(rt.postPhoto))

	// Comment Post
	rt.router.POST("/users/:user_id/posts/:post_id/comments", rt.wrap(rt.postComment))

	// Uncomment Post
	rt.router.DELETE("/users/:user_id/posts/:post_id/comments/:comment_id", rt.wrap(rt.deleteComment))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
