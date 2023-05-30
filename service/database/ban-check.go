package database

// Fuction that checks if the requesting user was banned by anotherone
func (db *appdbimpl) BanCheck(requestingUser UserId, targetUser UserId) (bool, error) {

	var count int
	var query = "SELECT COUNT(*) FROM banned_users WHERE banned = ? AND banner = ?"

	err := db.c.QueryRow(query, requestingUser.User_id, targetUser.User_id).Scan(&count)

	if err != nil {
		return true, err
	}

	if count == 1 {
		return true, nil
	}
	return false, nil
}
