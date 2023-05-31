package api

import (
	"time"

	"git.gabrielefabro.it/service/database"
)

// Error messages
const INTERNAL_ERROR_MSG = "internal server error"
const PNG_ERROR_MSG = "file is not a png format"
const JPG_ERROR_MSG = "file is not a jpg format"
const IMG_FORMAT_ERROR_MSG = "images must be jpeg or png"
const INVALID_JSON_ERROR_MSG = "invalid json format"
const INVALID_IDENTIFIER_ERROR_MSG = "identifier must be a string between 3 and 16 characters"

// JSON Error Structure
type JSONErrorMsg struct {
	Message string `json:"message"` // Error messages
}

// Profile struct represent a profile.
type Profile struct {
	User_id   string            `json:"user_id"`
	Username  string            `json:"username"`
	Posts     []database.Post   `json:"post"`
	Following []database.UserId `json:"following"`
	Follower  []database.UserId `json:"follower"`
}

// Post struct represent a post.
type Post struct {
	User_id          string             `json:"user_id"`
	Post_id          int64              `json:"post_id"`
	Publication_time time.Time          `json:"publication_time"`
	Like             []database.User    `json:"likes"`
	Comment          []database.Comment `json:"comments"`
}

// Comment struct represent a comment
type Comment struct {
	User_id    string `json:"user_id"`
	Post_id    int64  `json:"post_id"`
	Username   string `json:"username"`
	Text       string `json:"text"`
	Comment_id int64  `json:"comment_id"`
}

// PostId represent the id of profile
type PostId struct {
	Post_id int64 `json:"post_id"`
}

// PostId represent the id of profile
type CommentId struct {
	Comment_id int64 `json:"comment_id"`
}

// PostId represent the id of profile
type UserId struct {
	User_id string `json:"user_id"`
}

// Username represent the username of profile
type Username struct {
	Username string `json:"username"`
}

// TextComment represent the text of comment
type TextComment struct {
	TextComment string `json:"text"`
}

// Converts a Post_id from the api package to a Post_id of the database package
func (postId PostId) ToDatabase() database.PostId {
	return database.PostId{
		Post_id: postId.Post_id,
	}
}

// Converts an Username from the api package to a Username of the database package
func (username Username) ToDatabase() database.Username {
	return database.Username{
		Username: username.Username,
	}
}

// Converts a CommentId from the api package to a CommentId of the database package
func (comment CommentId) ToDatabase() database.CommentId {
	return database.CommentId{
		Comment_id: comment.Comment_id,
	}
}

// Converts a User from the api package to a User of the database package
func (userId UserId) ToDatabase() database.UserId {
	return database.UserId{
		User_id: userId.User_id,
	}
}

// Converts a TextComment from the api package to a TextComment of the database package
func (text TextComment) ToDatabase() database.TextComment {
	return database.TextComment{
		TextComment: text.TextComment,
	}
}

// Converts a Post from the api package to a Post of the database package
func (post Post) ToDatabase() database.Post {
	return database.Post{
		User_id:          post.User_id,
		Post_id:          post.Post_id,
		Publication_time: post.Publication_time,
		Like:             post.Like,
		Comment:          post.Comment,
	}
}

// Converts a Comment from the api package to a Comment of the database package
func (comment Comment) ToDatabase() database.Comment {
	return database.Comment{
		Comment_id: comment.Comment_id,
		Post_id:    comment.Post_id,
		Username:   comment.Username,
		User_id:    comment.User_id,
		Text:       comment.Text,
	}
}
