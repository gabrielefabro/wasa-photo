package database

// Function that modifies a user's username
func (db *appdbimpl) ModifyUserName(user User, newUserName Username) error {

	var query = "UPDATE users SET username = ? WHERE id = ?"

	_, err := db.c.Exec(query, newUserName.Username, user.User_id)
	if err != nil {
		return err
	}
	return nil
}
