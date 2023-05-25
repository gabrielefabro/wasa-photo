package database

// Function that return a profile that matches with the username passed as argoument
func (db *appdbimpl) GetUserProfile(searcher User, userToSearch User) (Profile, int64, error) {

	var query = "SELECT * FROM profiles WHERE ((user_id LIKE ?) OR (username LIKE ?)) AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)"

	rows, err := db.c.Query(query, userToSearch.User_id, userToSearch.User_id, searcher.User_id)
	if err != nil {
		return Profile{}, 0, err
	}
	defer func() { _ = rows.Close() }()

	var profile Profile
	err = rows.Scan(&profile.User_id, &profile.Posts, &profile.Following, &profile.Follower)
	if err != nil {
		return Profile{}, 0, err
	}

	if rows.Err() != nil {
		return Profile{}, 0, err
	}

	username, err := db.GetUserName(profile.User_id)
	if err != nil {
		return Profile{}, 0, err
	}
	profile.Username = username

	return profile, 1, nil
}
