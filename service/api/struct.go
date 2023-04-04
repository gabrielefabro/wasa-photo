package api

import (
	"time"
	"wasa-photo/service/database"
)

// Profile struct represent a profile for the APIs
type Profile struct {
	User            User   `json:"users"`
	Bio             string `json:"bio"`
	Posts           []Post `json:"posts"`
	Following_count uint64 `json:"following_count"`
	Follower_count  uint64 `json:"follower_count"`
}

// Post struct represent a post for the APIs
type Post struct {
	User             User      `json:"users"`
	Post_id          uint64    `json:"post_id"`
	Publication_time time.Time `json:"pubblication_time"`
	Bio              string    `json:"bio"`
	Photo_url        string    `json:"photo_url"`
	Like_count       uint64    `json:"likes"`
	Comment_count    uint64 `json:"comments"`
}

// Comment struct represent a comment for the APIs
type Comment struct {
	User         User      `json:"users"`
	Post_id      uint64    `json:"post_id"`
	Text         string    `json:"text"`
	Comment_id   uint64    `json:"comment_id"`
	Time_comment time.Time `json:"time_comment"`
}

// User represent the couple ID and UserName for the APIs
type User struct {
	User_id  uint64 `json:"user_id"`
	UserName string `json:"username"`
}

// PostId represent the id of profile for the APIs
type PostId struct {
	Post_id uint64 `json:"post_id"`
}

// PostId represent the id of profile for the APIs
type CommentId struct {
	Comment_id uint64 `json:"comment_id"`
}

// PostId represent the id of profile for the APIs
type UserId struct {
	User_id int64 `json:"user_id"`
}

// UserName represent the username of the profile for the APIs
type Username struct {
	Username string `json:"username"`
}

// TextComment represent a the text of a comment  for the APIs
type TextComment struct {
	TextComment string `json:"text"`
}

// Converts a User from the api package to a User of the database package
func (user User) ToDatabase() database.User {
	return database.User{
		User_id:  user.User_id,
		UserName: user.UserName,
	}
}

// Converts a Photo from the api package to a Photo of the database package
func (post Post) ToDatabase() database.Post {
	return database.Post{
		Post_id: post.Post_id,
		User_id: post.User.User_id,
		Publication_time: post.Publication_time,
		Bio: post.Bio,

}
