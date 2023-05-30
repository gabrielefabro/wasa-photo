package database

// Function that adds a comment of a user to a post
func (db *appdbimpl) CommentPost(postId PostId, UserId UserId, text TextComment) (int64, error) {

	var query = "INSERT INTO comments (post_id,user_id,text) VALUES (?, ?, ?)"

	res, err := db.c.Exec(query, postId.Post_id, UserId.User_id, text.TextComment)
	if err != nil {
		return -1, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return commentId, nil
}
