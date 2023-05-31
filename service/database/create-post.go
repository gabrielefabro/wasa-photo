package database

// Function that creates a post on the database
func (db *appdbimpl) CreatePost(post Post) (int64, error) {

	res, err := db.c.Exec("INSERT INTO posts (user_id,publication_time) VALUES (?,?)",
		post.User_id, post.Publication_time)

	if err != nil {
		return -1, err
	}

	postId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return postId, nil
}
