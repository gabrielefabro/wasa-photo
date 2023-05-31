package database

// Function that adds a new user in the Database
func (db *appdbimpl) CreateUser(userId UserId) error {

	var query = "INSERT INTO users (user_id,username) VALUES (?, ?)"

	_, err := db.c.Exec(query, userId.User_id, userId.User_id)

	if err != nil {
		return err
	}

	return nil
}
