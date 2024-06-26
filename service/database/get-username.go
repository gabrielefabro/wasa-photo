package database

// Function that gets a user's username
func (db *appdbimpl) GetUserName(userId UserId) (string, error) {

	var username string
	var query = "SELECT username FROM users WHERE user_id = ?"

	err := db.c.QueryRow(query, userId.User_id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
