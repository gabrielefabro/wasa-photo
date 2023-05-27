package database

// Database function that removes a photo from the database
func (db *appdbimpl) DeletePost(user User, postId PostId) error {

	_, err := db.c.Exec("DELETE FROM posts WHERE user_id = ? AND post_id = ? ",
		user.User_id, postId)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	return nil
}
