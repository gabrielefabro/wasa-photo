package database

// Function that removes a follower from a speficif user
func (db *appdbimpl) UnfollowUser(follower UserId, followed UserId) error {

	var query = "DELETE FROM followers WHERE(follower = ? AND followed = ?)"

	_, err := db.c.Exec(query, follower.User_id, followed.User_id)
	if err != nil {
		return err
	}

	return nil
}
