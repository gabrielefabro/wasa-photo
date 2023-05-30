package database

// Fuction that allow to an user to ban another user
func (db *appdbimpl) BanUser(banner UserId, banned UserId) error {

	var query = "INSERT INTO banned_users (banner,banned) VALUES (?, ?)"

	_, err := db.c.Exec(query, banner.User_id, banned.User_id)
	if err != nil {
		return err
	}

	return nil
}
