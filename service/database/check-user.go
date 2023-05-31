package database

// Function that check if an targetUser exists
func (db *appdbimpl) CheckUser(searchedUser UserId) (bool, error) {

	var count int
	var query = "SELECT COUNT(*) FROM users WHERE user_id = ?"

	err := db.c.QueryRow(query, searchedUser.User_id).Scan(&count)

	if err != nil {
		return true, err
	}

	if count == 1 {
		return true, nil
	}
	return false, nil
}
