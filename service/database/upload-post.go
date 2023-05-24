package database

// Function that creates a post on the database
func (db *appdbimpl) UploadPost(post Post) (int64, error) {

	var query = "INSERT INTO posts (user_id,publication_time,bio) VALUES (?,?,?)"

	res, err := db.c.Exec(query, post.User_id, post.Publication_time)

	if err != nil {
		// Error executing query
		return -1, err
	}

	postId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return postId, nil
}
