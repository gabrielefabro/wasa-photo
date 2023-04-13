package database

// Database function that return a profile that matches with the username passed as argoument
func (db *appdbimpl) GetProfile(searcher User, userToSearch User) (Profile, int64, error) {

	rows, err := db.c.Query("SELECT profile FROM profiles WHERE user_id = ? AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		userToSearch.User_id, userToSearch.User_id, searcher.User_id)
	if err != nil {
		return Profile{}, 0, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	var profile Profile
	err = rows.Scan(&profile.User, &profile.Posts, &profile.Bio, &profile.Following, &profile.Follower)
	if err != nil {
		return Profile{}, 0, err
	}

	if rows.Err() != nil {
		return Profile{}, 0, err
	}

	return profile, 1, nil
}

// Database function that filters the users by a parameter. Any partial match is included in the result.
// Returns a list of matching users (either by nickname or identifier)
func (db *appdbimpl) SearchUser(searcher User, userToSearch User) ([]User, error) {

	rows, err := db.c.Query("SELECT * FROM users WHERE ((user_id LIKE ?) OR (username LIKE ?)) AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		userToSearch.User_id, userToSearch.User_id, searcher.User_id)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset.
	var res []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.User_id, &user.UserName)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}
