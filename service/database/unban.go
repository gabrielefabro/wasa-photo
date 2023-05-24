package database

// Fuction that removes an user from the banned list of another one
func (db *appdbimpl) UnbanUser(banner User, banned User) error {

	var query = "DELETE FROM banned_users WHERE (banner = ? AND banned = ?)"

	_, err := db.c.Exec(query, banner.User_id, banned.User_id)
	if err != nil {
		return err
	}

	return nil
}
