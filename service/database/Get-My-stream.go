package database

import "fmt"

func (db *appdbimpl) GetMyStream(id uint64) ([]Post, error) {
	rows, err := db.c.Query(`
        SELECT posts.profile_id, posts.timestamp, posts.caption, posts.like_count, posts.comment_count, posts.image_url
        FROM posts
        JOIN follows ON posts.profile_id = follows.followed_profile_id
        WHERE follows.follower_user_id = ?
        ORDER BY posts.timestamp DESC
    `, id)
	if err != nil {
		return nil, fmt.Errorf("error querying posts: %v", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.User, &post.PublicationTime, &post.Bio, &post.LikeCount, &post.CommentCount, &post.PhotoUrl)
		if err != nil {
			return nil, fmt.Errorf("error scanning post: %v", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return posts, nil
}
