package database

// Function that check if an targetUser exists
func (db *appdbimpl) CheckUser(targetUser User) (bool, error) {

	var count int
	var query = "SELECT COUNT(*) FROM users WHERE user_id = ?"

	err := db.c.QueryRow(query, targetUser.User_id).Scan(&count)

	if err != nil {
		return true, err
	}

	// If the counter is 1 then the user exists
	if count == 1 {
		return true, nil
	}
	return false, nil
}
