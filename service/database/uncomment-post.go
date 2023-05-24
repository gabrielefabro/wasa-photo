package database

// Function that removes a comment of a user from a photo
func (db *appdbimpl) UncommentPost(postId PostId, user User, comment CommentId) error {

	var query = "DELETE FROM comments WHERE (post_id = ? AND user_id = ? AND comment_id = ?)"

	_, err := db.c.Exec(query, postId, user.User_id, comment.Comment_id)
	if err != nil {
		return err
	}

	return nil
}
