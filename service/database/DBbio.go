package database

// Database function that add a bio to a profile that exists yet
func (db *appdbimpl) SetBio(userID int, bio string) error {
	_, err := db.c.Exec("UPDATE profiles SET bio=? WHERE user_id=?", bio, userID)
	if err != nil {
		return err
	}
	return nil
}
