package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrProfileDoesNotExist = errors.New("profile does not exist")
var ErrUserBanned = errors.New("user is banned")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// CreateUser create a new user
	CreateUser(UserId) error

	// GetUserName get the username of a user
	GetUserName(UserId) (string, error)

	// ModifyUserName set a new Username for an existing profile
	ChangeUserName(UserId, Username) error

	// GetMyStream returns the stream of the id passed as argoument
	GetMyStream(UserId) ([]Post, error)

	// GetFollowers returns the followers list
	GetFollowers(UserId) ([]UserId, error)

	// GetFollowings returns the followings list
	GetFollowings(UserId) ([]UserId, error)

	// GetPosts return all the post from one profile
	GetPosts(requestingUser UserId, targetUser UserId) ([]Post, error)

	// GetComments return all the comments from one post
	GetComments(requestingUser UserId, requestedUser UserId, postId PostId) ([]Comment, error)

	// GetLikes return all the like from on post
	GetLikes(requestingUser UserId, requestedUser UserId, postId PostId) ([]User, error)

	// FollowUser adds one profile from the followers list
	FollowUser(followed UserId, follower UserId) error

	// UnfollowUser removes one profile from the followers list
	UnfollowUser(followed UserId, follower UserId) error

	// BanUser adds one profile from the bans list
	BanUser(banned UserId, banner UserId) error

	// UnbanUser remove one profile from the bans list
	UnbanUser(banned UserId, banner UserId) error

	// LikePost add a like to the likes list
	LikePost(PostId, UserId) error

	// UnlikePost removes a like to the Unlikes list
	UnlikePost(PostId, UserId) error

	// CommentPost adds a comment in the comments list
	CommentPost(PostId, UserId, TextComment) (int64, error)

	// UncommentPost adds a comment in the comments list
	UncommentPost(PostId, UserId, CommentId) error

	// DeletePost deletes a post by his id
	DeletePost(UserId, PostId) error

	// CreatePost create a new post in the database
	CreatePost(Post) (int64, error)

	// BannedCheck control if an user is banned by anotherone
	BanCheck(banned UserId, banner UserId) (bool, error)

	// CheckUser control if an user exist
	CheckUser(UserId) (bool, error)

	// SearchUser searches all the users that match the given name
	SearchUser(searcher UserId, userToSearch UserId) ([]User, error)

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Activate foreign keys for db

	_, errPramga := db.Exec(`PRAGMA foreign_keys= ON`)
	if errPramga != nil {
		return nil, fmt.Errorf("error setting pragmas: %w", errPramga)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createDatabase(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

// Creates all the necessary sql tables for the WASAPhoto app.
func createDatabase(db *sql.DB) error {
	tables := [6]string{
		`CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(15) NOT NULL PRIMARY KEY,
			username VARCHAR(15) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id VARCHAR(15) NOT NULL,
			publication_time DATETIME NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS  likes (
			post_id INTEGER NOT NULL,
			user_id VARCHAR(15) NOT NULL,
			PRIMARY KEY (post_id,user_id),
			FOREIGN KEY(post_id) REFERENCES posts (post_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS comments (
			comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id VARCHAR(15) NOT NULL,
			post_id INTEGER NOT NULL,
			text VARCHAR(50) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
			FOREIGN KEY(post_id) REFERENCES posts (post_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS banned_users (
			banner VARCHAR(15) NOT NULL,
			banned VARCHAR(15) NOT NULL,
			PRIMARY KEY (banner,banned),
			FOREIGN KEY(banner) REFERENCES users (user_id) ON DELETE CASCADE,
			FOREIGN KEY(banned) REFERENCES users (user_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS followers(
			follower VARCHAR(15) NOT NULL,
			followed VARCHAR(15) NOT NULL,
			PRIMARY KEY (follower,followed),
			FOREIGN KEY(follower) REFERENCES users (user_id) ON DELETE CASCADE,
			FOREIGN KEY(followed) REFERENCES users (user_id) ON DELETE CASCADE
			);`,
	}

	// Iteration to create all the needed sql schemas
	for i := 0; i < len(tables); i++ {

		sqlStmt := tables[i]
		_, err := db.Exec(sqlStmt)

		if err != nil {
			return err
		}
	}
	return nil
}
