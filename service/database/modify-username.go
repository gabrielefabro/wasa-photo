package database

// Function that modifies a user's username
func (db *appdbimpl) ModifyUserName(userId UserId, newUserName Username) error {

	var query = "UPDATE users SET username = ? WHERE user_id = ?"

	_, err := db.c.Exec(query, newUserName.Username, userId.User_id)
	if err != nil {
		return err
	}
	return nil
}
