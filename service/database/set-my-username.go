package database

func (db *appdbimpl) SetMyUserName(u User) error {
	res, err := db.c.Exec("UPDATE profile SET username=?", u.UserName)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ProfileDoesNotExist
	}
	return nil
}
