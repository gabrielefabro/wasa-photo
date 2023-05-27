package database

// Function that removes a like of a user from a photo
func (db *appdbimpl) UnlikePost(postId PostId, user User) error {

	var query = "DELETE FROM likes WHERE(post_id = ? AND user_id = ?)"

	_, err := db.c.Exec(query, postId.Post_id, user.User_id)
	if err != nil {
		return err
	}

	return nil
}
