package database

// Database function that gets a user's nickname
func (db *appdbimpl) GetUserName(user User) (string, error) {

	var username string

	err := db.c.QueryRow(`SELECT username FROM users WHERE id = ?`, user.User_id).Scan(&username)
	if err != nil {
		// Error during the execution of the query
		return "", err
	}
	return username, nil
}

// Database function that modifies a user's nickname
func (db *appdbimpl) ModifyUserName(user User, newUserName Username) error {

	_, err := db.c.Exec(`UPDATE users SET username = ? WHERE id = ?`, newUserName.Username, user.User_id)
	if err != nil {
		// Error during the execution of the query
		return err
	}
	return nil
}
