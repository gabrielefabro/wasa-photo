package database

// Function that adds a like of a user to a photo
func (db *appdbimpl) LikePost(postId PostId, user User) error {

	var query = "INSERT INTO likes (post_id,user_id,username) VALUES (?, ?, ?)"

	_, err := db.c.Exec(query, postId, user.User_id, user.Username)
	if err != nil {
		return err
	}

	return nil
}
