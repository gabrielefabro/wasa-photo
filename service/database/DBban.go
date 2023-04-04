package database

// Database fuction that get the permission to an user to ban another one
func (db *appdbimpl) BanUser(banner User, banned User) error {

	_, err := db.c.Exec("INSERT INTO banned_users (banner,banned) VALUES (?, ?)", banner.User_id, banned.User_id)
	if err != nil {
		return err
	}

	return nil
}

// Database fuction that removes an user from the banned list of another one
func (db *appdbimpl) UnbanUser(banner User, banned User) error {

	_, err := db.c.Exec("DELETE FROM banned_users WHERE (banner = ? AND banned = ?)", banner.User_id, banned.User_id)
	if err != nil {
		return err
	}

	return nil
}

// Database fuction that checks if the requesting user was banned by anotherone
func (db *appdbimpl) BanCheck(requestingUser User, targetUser User) (bool, error) {

	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM banned_users WHERE banned = ? AND banner = ?",
		requestingUser.User_id, targetUser.User_id).Scan(&cnt)

	if err != nil {
		// Count always returns a row thanks to COUNT(*), so this situation should not happen
		return true, err
	}

	// If the counter is 1 then the user was banned
	if cnt == 1 {
		return true, nil
	}
	return false, nil
}
