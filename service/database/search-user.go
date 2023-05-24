package database // Function that filters the users by a parameter
func (db *appdbimpl) SearchUser(searcher User, userToSearch User) ([]User, error) {

	var query = "SELECT * FROM users WHERE ((user_id LIKE ?) OR (username LIKE ?)) AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)"

	rows, err := db.c.Query(query, userToSearch.User_id, userToSearch.User_id, searcher.User_id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var res []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.User_id, &user.Username)
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
