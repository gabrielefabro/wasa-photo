package database

// Function that return the list of followers of a user
func (db *appdbimpl) GetFollowers(requestinUser User) ([]User, error) {

	var query = "SELECT follower FROM followers WHERE followed = ?"

	rows, err := db.c.Query(query, requestinUser.User_id)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var followers []User
	for rows.Next() {
		var follower User
		err = rows.Scan(&follower.User_id)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return followers, nil
}