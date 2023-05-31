package database

// Function that delete a post from the database
func (db *appdbimpl) DeletePost(userId UserId, postId PostId) error {

	_, err := db.c.Exec("DELETE FROM posts WHERE user_id = ? AND post_id = ? ",
		userId.User_id, postId.Post_id)
	if err != nil {
		return err
	}

	return nil
}
