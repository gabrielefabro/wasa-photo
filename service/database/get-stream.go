package database

// Function that gets the stream of a user (photos of people that are followed by the latter)
func (db *appdbimpl) GetStream(user User) ([]Post, error) {

	var query = "SELECT * FROM posts WHERE user_id IN (SELECT followed FROM followers WHERE follower = ?) ORDER BY publication_time DESC"

	rows, err := db.c.Query(query, user.User_id)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var res []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Post_id, &post.User_id, &post.Publication_time)
		if err != nil {
			return nil, err
		}

		username, err := db.GetUserName(post.User_id)
		if err != nil {
			return nil, err
		}
		post.Username = username

		res = append(res, post)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}
