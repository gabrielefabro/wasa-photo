package database

// Function that adds a new user in the database
func (db *appdbimpl) CreateUser(user User) error {

	var query = "INSERT INTO users (user_id,username) VALUES (?, ?)"

	_, err := db.c.Exec(query, user.User_id, user.Username)

	if err != nil {
		return err
	}

	return nil
}
