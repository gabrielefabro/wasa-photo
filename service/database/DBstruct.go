package database

import "time"

// Profile struct represent a profile.
type Profile struct {
	User      User   `json:"users"`
	Posts     []Post `json:"posts"`
	Following []User `json:"following"`
	Follower  []User `json:"follower"`
}

// Post struct represent a post.
type Post struct {
	User_id          string    `json:"users"`
	Post_id          uint64    `json:"post_id"`
	Publication_time time.Time `json:"pubblication_time"`
	Photo_url        string    `json:"photo_url"`
	Like             []User    `json:"likes"`
	Comment          []Comment `json:"comments"`
}

// Comment struct represent a comment
type Comment struct {
	User         User      `json:"users"`
	Post_id      uint64    `json:"post_id"`
	Text         string    `json:"text"`
	Comment_id   uint64    `json:"comment_id"`
	Time_comment time.Time `json:"time_comment"`
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
