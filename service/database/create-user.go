package database

// Function that adds a new user in the database
func (db *appdbimpl) CreateUser(user_id string) error {

	var query = "INSERT INTO users (user_id,username) VALUES (?, ?)"

	_, err := db.c.Exec(query, user_id, user_id)

	if err != nil {
		return err
	}

	return nil
}
