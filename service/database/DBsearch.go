package database

// Database function that return a profile that matches with the username passed as argoument
func (db *appdbimpl) GetProfile(searcher User, userToSearch User) (Profile, error) {

	rows, err := db.c.Query("SELECT profile FROM profiles WHERE user_id = ? AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		userToSearch.User_id, userToSearch.User_id, searcher.User_id)
	if err != nil {
		return Profile{}, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	var profile Profile
	err = rows.Scan(&profile.User, &profile.Posts, &profile.Bio, &profile.Following, &profile.Follower)
	if err != nil {
		return Profile{}, err
	}

	if rows.Err() != nil {
		return Profile{}, err
	}

	return profile, nil
}
