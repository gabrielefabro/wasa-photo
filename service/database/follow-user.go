package database

// Database function that adds a follower to a user
func (db *appdbimpl) FollowUser(follower User, followed User) error {

	_, err := db.c.Exec("INSERT INTO followers (follower,followed) VALUES (?, ?)",
		follower.User_id, followed.User_id)
	if err != nil {
		return err
	}

	return nil
}
