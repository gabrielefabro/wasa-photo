package database

// Function that retrieves the list of users that liked a photo
func (db *appdbimpl) GetLikes(requestingUser UserId, requestedUser UserId, postId PostId) ([]User, error) {

	var query = "SELECT user_id FROM likes WHERE post_id = ? AND user_id NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) " +
		"AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)"

	rows, err := db.c.Query(query, postId.Post_id, requestingUser.User_id, requestedUser.User_id, requestingUser.User_id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var likes []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.User_id)
		if err != nil {
			return nil, err
		}

		username, err := db.GetUserName(UserId{User_id: user.User_id})
		if err != nil {
			return nil, err
		}
		user.Username = username

		likes = append(likes, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return likes, nil
}
