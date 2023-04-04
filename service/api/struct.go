package api

import (
	"time"
	"wasa-photo/service/database"
)

// Profile struct represent a profile.
type Profile struct {
	User      User            `json:"users"`
	Bio       string          `json:"bio"`
	Posts     []database.Post `json:"posts"`
	Following []database.User `json:"following"`
	Follower  []database.User `json:"follower"`
}

// Post struct represent a post.
type Post struct {
	User             database.User      `json:"users"`
	Post_id          uint64             `json:"post_id"`
	Publication_time time.Time          `json:"pubblication_time"`
	Bio              string             `json:"bio"`
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
	User_id  uint64 `json:"user_id"`
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
	User_id int64 `json:"user_id"`
}

type Username struct {
	Username string `json:"username"`
}

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

// Converts a Post from the api package to a Post of the database package
func (post Post) ToDatabase() database.Post {
	return database.Post{
		User:             post.User,
		Publication_time: post.Publication_time,
		Bio:              post.Bio,
		Like:             post.Like,
		Comment:          post.Comment,
	}
}

func (comment Comment) ToDatabase() database.Comment {
	return database.Comment{
		Comment_id:   comment.Comment_id,
		Post_id:      comment.Post_id,
		User:         comment.User,
		Text:         comment.Text,
		Time_comment: comment.Time_comment,
	}
}
