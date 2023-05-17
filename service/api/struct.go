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
	User      User            `json:"users"`
	Posts     []database.Post `json:"posts"`
	Following []database.User `json:"following"`
	Follower  []database.User `json:"follower"`
}

// Post struct represent a post.
type Post struct {
	User_id          string             `json:"users"`
	Post_id          uint64             `json:"post_id"`
	Publication_time time.Time          `json:"pubblication_time"`
	Photo_url        string             `json:"photo_url"`
	Like             []database.User    `json:"likes"`
	Comment          []database.Comment `json:"comments"`
}

// Comment struct represent a comment
type Comment struct {
	User         database.User `json:"users"`
	Post_id      uint64        `json:"post_id"`
	Text         string        `json:"text"`
	Comment_id   uint64        `json:"comment_id"`
	Time_comment time.Time     `json:"time_comment"`
}

// User represent the couple ID and UserName
type User struct {
	User_id  string `json:"user_id"`
	UserName string `json:"username"`
}

// PostId represent the id of profile
type PostId struct {
	Post_id uint64 `json:"post_id"`
}

// PostId represent the id of profile
type CommentId struct {
	Comment_id uint64 `json:"comment_id"`
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
func (post_id PostId) ToDatabase() database.PostId {
	return database.PostId{
		Post_id: post_id.Post_id,
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
func (user User) ToDatabase() database.User {
	return database.User{
		User_id:  user.User_id,
		UserName: user.UserName,
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
		Publication_time: post.Publication_time,
		Like:             post.Like,
		Comment:          post.Comment,
	}
}

// Converts a Comment from the api package to a Comment of the database package
func (comment Comment) ToDatabase() database.Comment {
	return database.Comment{
		Comment_id:   comment.Comment_id,
		Post_id:      comment.Post_id,
		User:         comment.User,
		Text:         comment.Text,
		Time_comment: comment.Time_comment,
	}
}
