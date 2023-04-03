package database

// Database function that gets the stream of a user (photos of people that are followed by the latter)
func (db *appdbimpl) GetMyStream(user User) ([]Post, error) {

	rows, err := db.c.Query(`SELECT * FROM posts WHERE user_id IN (SELECT followed FROM followers WHERE follower = ?) ORDER BY publication_time DESC`,
		user.User_id)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset
	var res []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Post_id, &post.User, &post.Publication_time) //  &photo.Comments, &photo.Likes,
		if err != nil {
			return nil, err
		}
		res = append(res, post)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}

// Database function that adds a new user in the database upon registration
func (db *appdbimpl) CreateUser(user User) error {

	_, err := db.c.Exec("INSERT INTO users (user_id,username) VALUES (?, ?)",
		user.User_id, user.UserName)

	if err != nil {
		return err
	}

	return nil
}
