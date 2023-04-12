package database

// Database function that add a bio to a profile that exists yet
func (db *appdbimpl) SetBio(userID int, bio string) error {
	_, err := db.c.Exec("UPDATE profiles SET bio=? WHERE user_id=?", bio, userID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetBio(user User) (string, error) {
	var bio string
	err := db.c.QueryRow("SELECT bio FROM profiles WHERE user_id = ?", user.User_id).Scan(bio)
	if err != nil {
		return bio, ErrUserBanned
	}

	return bio, nil
}
