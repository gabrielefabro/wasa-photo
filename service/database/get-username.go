package database

// Function that gets a user's nickname
func (db *appdbimpl) GetUserName(user_id string) (string, error) {

	var username string
	var query = "SELECT username FROM users WHERE id = ?"

	err := db.c.QueryRow(query, user_id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
