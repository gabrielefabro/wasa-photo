package database

// Function that removes a like of a user from a post
func (db *appdbimpl) UnlikePost(postId PostId, userId UserId) error {

	var query = "DELETE FROM likes WHERE(post_id = ? AND user_id = ?)"

	_, err := db.c.Exec(query, postId.Post_id, userId.User_id)
	if err != nil {
		return err
	}

	return nil
}
