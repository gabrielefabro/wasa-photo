package database

// Database function that retrieves the list of comments of a photo (minus the comments from users that banned the requesting user)
func (db *appdbimpl) GetComments(requestingUser User, requestedUser User, Post Post) ([]Comment, error) {

	rows, err := db.c.Query("SELECT * FROM comments WHERE post_id = ? AND user_id NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
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
		username, err := db.GetUserName(comment.User)
		if err != nil {
			return nil, err
		}
		comment.User.UserName = username

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}

// Database function that adds a comment of a user to a photo
func (db *appdbimpl) CommentPhoto(post_id PostId, user User, text TextComment) (int64, error) {

	res, err := db.c.Exec("INSERT INTO comments (post_id,username,user_id,text) VALUES (?, ?, ?)",
		post_id, user.UserName, user.User_id, text)
	if err != nil {
		// Error executing query
		return -1, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		// Error getting id returned by last db operation (commentId)
		return -1, err
	}

	return commentId, nil
}

// Database function that removes a comment of a user from a photo
func (db *appdbimpl) UncommentPhoto(post_id PostId, user User, comment CommentId) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE (post_id = ? AND user_id = ? AND comment_id = ?)",
		post_id, user.User_id, comment.Comment_id)
	if err != nil {
		return err
	}

	return nil
}
