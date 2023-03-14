package database

// Database function that retrieves the list of comments of a photo (minus the comments from users that banned the requesting user)
func (db *appdbimpl) GetCompleteCommentsList(requestingUser User, requestedUser User, Post Post) ([]Comment, error) {

	rows, err := db.c.Query("SELECT * FROM comments WHERE poat_id = ? AND user_id NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
		"AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		Post.Post_id, requestingUser.User_id, requestedUser.User_id, requestingUser.User_id)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the comments in the resulset (comments of the photo with authors that didn't ban the requesting user).
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.User.User_id, &comment.Post_id, &comment.User.UserName, &comment.Text, &comment.Comment_id, &comment.Time_comment)
		if err != nil {
			return nil, err
		}

		// Get the nickname of the user that commented
		nickname, err := db.GetNickname(User{User_id: comment.User.User_id})
		if err != nil {
			return nil, err
		}
		comment.Nickname = nickname

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}
