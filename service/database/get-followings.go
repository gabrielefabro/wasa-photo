package database

// Function that return the list of users followed by the user
func (db *appdbimpl) GetFollowings(requestinUser UserId) ([]UserId, error) {

	var query = "SELECT followed FROM followers WHERE follower = ?"

	rows, err := db.c.Query(query, requestinUser.User_id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var following []UserId
	for rows.Next() {
		var followed UserId
		err = rows.Scan(&followed.User_id)
		if err != nil {
			return nil, err
		}
		following = append(following, followed)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return following, nil
}
