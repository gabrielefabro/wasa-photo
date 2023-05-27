package database

// Function that adds a comment of a user to a post
func (db *appdbimpl) CommentPost(postId PostId, user User, text TextComment) (int64, error) {

	var query = "INSERT INTO comments (post_id,username,user_id,text) VALUES (?, ?, ?, ?)"

	res, err := db.c.Exec(query, postId.Post_id, user.Username, user.User_id, text)
	if err != nil {
		// Error executing query
		return -1, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return commentId, nil
}
