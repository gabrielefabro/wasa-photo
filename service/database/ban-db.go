package database

// Database fuction that allows a user (banner) to ban another one (banned)
func (db *appdbimpl) BanUser(banner User, banned User) error {

	_, err := db.c.Exec("INSERT INTO banned_users (banner,banned) VALUES (?, ?)", banner.User_id, banned.User_id)
	if err != nil {
		return err
	}

	return nil
}

// Database fuction that removes a user (banned) from the banned list of another one (banner)
func (db *appdbimpl) UnbanUser(banner User, banned User) error {

	_, err := db.c.Exec("DELETE FROM banned_users WHERE (banner = ? AND banned = ?)", banner.User_id, banned.User_id)
	if err != nil {
		return err
	}

	return nil
}

// Database function that checks if a user exists
func (db *appdbimpl) BanCheck(targetUser User) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE user_id = ?",
		targetUser.User_id).Scan(&count)

	if err != nil {
		// Count always returns a row thanks to COUNT(*), so this situation should not happen
		return true, err
	}

	// If the counter is 1 then the user exists
	if count == 1 {
		return true, nil
	}
	return false, nil
}
