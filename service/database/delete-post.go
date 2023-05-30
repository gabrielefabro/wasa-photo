package database

// Function that delete a post from the database
func (db *appdbimpl) DeletePost(UserId UserId, postId PostId) error {

	_, err := db.c.Exec("DELETE FROM posts WHERE user_id = ? AND post_id = ? ",
		UserId.User_id, postId.Post_id)
	if err != nil {
		return err
	}

	return nil
}
