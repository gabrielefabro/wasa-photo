package database

func (db *appdbimpl) SetMyUserName(u User) error {

	query = "UPDATE profile SET username = ?"

	res, err := db.c.Exec(query, u.UserName)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		return err
	} else if affected == 0 {
		
		// If we didn't delete any row, then the profile didn't exist
		return ProfileDoesNotExist
	}

	return nil
}
