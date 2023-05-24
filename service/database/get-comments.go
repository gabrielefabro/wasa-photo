package database

// Function that return the list of comments of a post
func (db *appdbimpl) GetComments(requestingUser User, requestedUser User, post Post) ([]Comment, error) {

	var query = "SELECT * FROM comments WHERE post_id = ? AND user_id NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) " +
		"AND user_id NOT IN (SELECT banner FROM banned_users WHERE banned = ?)"

	rows, err := db.c.Query(query, post.Post_id, requestingUser.User_id, requestedUser.User_id, requestingUser.User_id)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.User_id, &comment.Post_id, &comment.Text, &comment.Comment_id, &comment.Time_comment)
		if err != nil {
			return nil, err
		}

		username, err := db.GetUserName(comment.User_id)
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
