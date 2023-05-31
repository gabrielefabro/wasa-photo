package database

// Function that return a requested post
func (db *appdbimpl) GetPhoto(requestinUser UserId, targetPost PostId) (Post, error) {

	var post Post
	var query = "SELECT * FROM posts WHERE (post_id = ?) AND user_id NOT IN (SELECT banner FROM banned_user WHERE banned = ?)"

	err := db.c.QueryRow(query, targetPost.Post_id, requestinUser.User_id).Scan(&post)

	if err != nil {
		return Post{}, ErrUserBanned
	}

	return post, nil

}
