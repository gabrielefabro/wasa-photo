package database

// Function that adds a like of a user to a post
func (db *appdbimpl) LikePost(postId PostId, UserId UserId) error {

	var query = "INSERT INTO likes (post_id,user_id) VALUES (?, ?)"

	_, err := db.c.Exec(query, postId.Post_id, UserId.User_id)
	if err != nil {
		return err
	}

	return nil
}
