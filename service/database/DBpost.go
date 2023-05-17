package database

// Database function that retrieves the list of photos of a user (only if the requesting user is not banned by that user)
func (db *appdbimpl) GetPosts(requestingUser User, targetUser User) ([]Post, error) { // requestinUser User,

	rows, err := db.c.Query("SELECT * FROM posts WHERE user_id = ? ORDER BY publication_time DESC", targetUser.User_id)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the photos in the resulset
	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Post_id, &post.User_id, &post.Publication_time, &post.Photo_url)
		if err != nil {
			return nil, err
		}

		comments, err := db.GetComments(requestingUser, targetUser, post)
		if err != nil {
			return nil, err
		}
		post.Comment = comments

		likes, err := db.GetLikes(requestingUser, targetUser, post)
		if err != nil {
			return nil, err
		}
		post.Like = likes

		posts = append(posts, post)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return posts, nil
}

// Database function that retrieves a specific photo (only if the requesting user is not banned by that owner of that photo).
func (db *appdbimpl) GetPhoto(requestinUser User, targetPhoto PostId) (Post, error) {

	var post Post
	err := db.c.QueryRow("SELECT * FROM posts WHERE (post_id = ?) AND user_id NOT IN (SELECT banner FROM banned_user WHERE banned = ?)",
		targetPhoto.Post_id, requestinUser.User_id).Scan(&post)

	if err != nil {
		return Post{}, ErrUserBanned
	}

	return post, nil

}

// Database function that creates a photo on the database and returns the unique photo id
func (db *appdbimpl) UploadPhoto(post Post) (int64, error) {

	res, err := db.c.Exec("INSERT INTO posts (user_id,publication_time,bio) VALUES (?,?,?)",
		post.User_id, post.Publication_time)

	if err != nil {
		// Error executing query
		return -1, err
	}

	postId, err := res.LastInsertId()
	if err != nil {
		// Error getting id returned by last db operation (photoId)
		return -1, err
	}

	return postId, nil
}

/*
Adding the owner is an additional security measure to delete photos that are actually owned
by that user
*/

// Database function that removes a photo from the database
func (db *appdbimpl) DeletePhoto(user User, postId PostId) error {

	_, err := db.c.Exec("DELETE FROM photos WHERE user_id = ? AND post_id = ? ",
		user.User_id, postId)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	return nil
}
