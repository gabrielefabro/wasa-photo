package database

// Function that return the list of comments of a post
func (db *appdbimpl) GetComments(requestingUser UserId, requestedUser UserId, postId PostId) ([]Comment, error) {

	var query = "SELECT * FROM comments WHERE post_id = ? AND user_id NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) " +
		"AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)"

	rows, err := db.c.Query(query, postId.Post_id, requestingUser.User_id, requestedUser.User_id, requestingUser.User_id)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		var user User
		err = rows.Scan(&comment.Comment_id, &comment.User_id, &comment.Post_id, &comment.Text)
		if err != nil {
			return nil, err
		}

		username, err := db.GetUserName(UserId{User_id: user.User_id})
		if err != nil {
			return nil, err
		}
		comment.Username = username

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}
