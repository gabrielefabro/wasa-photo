package database

// Fuction that removes an user from the banned list of anotherone
func (db *appdbimpl) UnbanUser(banner UserId, banned UserId) error {

	var query = "DELETE FROM banned_users WHERE (banner = ? AND banned = ?)"

	_, err := db.c.Exec(query, banner.User_id, banned.User_id)
	if err != nil {
		return err
	}

	return nil
}
