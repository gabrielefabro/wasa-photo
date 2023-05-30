package database

// Function that removes a comment of a user from a photo
func (db *appdbimpl) UncommentPost(postId PostId, user UserId, commentId CommentId) error {

	var query = "DELETE FROM comments WHERE (post_id = ? AND user_id = ? AND comment_id = ?)"

	_, err := db.c.Exec(query, postId.Post_id, user.User_id, commentId.Comment_id)
	if err != nil {
		return err
	}

	return nil
}
