package database

// Database function that retrieves the list of users that liked a photo
func (db *appdbimpl) GetLikesList(requestingUser User, requestedUser User, post Post) ([]User, error) {

	rows, err := db.c.Query("SELECT user_id FROM likes WHERE post_id = ? AND user_id NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
		"AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		post.Post_id, requestingUser.User_id, requestedUser.User_id, requestingUser.User_id)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users that liked the photo that didn't ban the requesting user).
	var likes []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.User_id)
		if err != nil {
			return nil, err
		}

		// Get the nickname of the user that liked the photo
		username, err := db.GetUserName(user)
		if err != nil {
			return nil, err
		}
		user.UserName = username

		likes = append(likes, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return likes, nil
}

// Database function that adds a like of a user to a photo
func (db *appdbimpl) LikePhoto(post Post, user User) error {

	_, err := db.c.Exec("INSERT INTO likes (post_id,user_id) VALUES (?, ?)", post.Post_id, user.User_id)
	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a like of a user from a photo
func (db *appdbimpl) UnlikePhoto(post Post, user User) error {

	_, err := db.c.Exec("DELETE FROM likes WHERE(id_photo = ? AND id_user = ?)", post.Post_id, user.User_id)
	if err != nil {
		return err
	}

	return nil
}
