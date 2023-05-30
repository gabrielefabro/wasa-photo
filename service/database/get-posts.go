package database

// Function that return the list of post of a user
func (db *appdbimpl) GetPosts(requestingUser UserId, targetUser UserId) ([]Post, error) { // requestinUser User,

	var query = "SELECT * FROM posts WHERE user_id = ? ORDER BY publication_time DESC"

	rows, err := db.c.Query(query, targetUser.User_id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Post_id, &post.User_id, &post.Publication_time)
		if err != nil {
			return nil, err
		}

		comments, err := db.GetComments(requestingUser, targetUser, PostId{Post_id: int64(post.Post_id)})
		if err != nil {
			return nil, err
		}
		post.Comment = comments

		likes, err := db.GetLikes(requestingUser, targetUser, PostId{Post_id: int64(post.Post_id)})
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
