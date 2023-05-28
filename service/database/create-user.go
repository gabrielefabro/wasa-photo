package database

// Function that adds a new user in the database
func (db *appdbimpl) CreateUser(userid userId) error {

	var query = "INSERT INTO users (user_id,username) VALUES (?, ?)"

	_, err := db.c.Exec(query, userid.User_id, userid.User_id)

	if err != nil {
		return err
	}

	return nil
}
