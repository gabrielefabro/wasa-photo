package database

import "time"

// Profile struct represent a profile.
type Profile struct {
	User_id   string   `json:"user_id"`
	Username  string   `json:"username"`
	Posts     []Post   `json:"post"`
	Following []UserId `json:"following"`
	Follower  []UserId `json:"follower"`
}

// Post struct represent a post.
type Post struct {
	User_id          string    `json:"user_id"`
	Post_id          int64     `json:"post_id"`
	Publication_time time.Time `json:"publication_time"`
	Like             []User    `json:"like"`
	Comment          []Comment `json:"comment"`
}

// Comment struct represent a comment
type Comment struct {
	User_id    string `json:"user_id"`
	Username   string `json:"username"`
	Post_id    int64  `json:"post_id"`
	Text       string `json:"text"`
	Comment_id int64  `json:"comment_id"`
}

// User represent the couple ID and UserName
type User struct {
	User_id  string `json:"user_id"`
	Username string `json:"username"`
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
