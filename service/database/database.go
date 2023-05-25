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

	// getUserProfile return the profile from the id passed as argument
	GetUserProfile(a User, b User) (Profile, int64, error)

	// CreateUser create a new user
	CreateUser(User) error

	// GetUserName get the username of a user
	GetUserName(user_id string) (string, error)

	// ModifyUserName set a new Username for an existing profile
	ModifyUserName(User, Username) error

	// GetMyStream returns the stream of the id passed as argoument
	GetMyStream(User) ([]Post, error)

	// GetFollowers returns the followers list
	GetFollowers(User) ([]User, error)

	// GetFollowings returns the followings list
	GetFollowings(User) ([]User, error)

	// GetPosts return all the post from one profile
	GetPosts(a User, b User) ([]Post, error)

	// GetPhoto return a single post from a profile
	GetPhoto(User, PostId) (Post, error)

	// FollowUser adds one profile from the followers list
	FollowUser(a User, b User) error

	// UnfollowUser removes one profile from the followers list
	UnfollowUser(a User, b User) error

	// BanUser adds one profile from the bans list
	BanUser(a User, b User) error

	// UnbanUser remove one profile from the bans list
	UnbanUser(a User, b User) error

	// LikePost add a like to the likes list
	LikePost(PostId, User) error

	// UnlikePost removes a like to the Unlikes list
	UnlikePost(PostId, User) error

	// CommentPost adds a comment in the comments list
	CommentPost(PostId, User, TextComment) (int64, error)

	// UncommentPost adds a comment in the comments list
	UncommentPost(PostId, User, CommentId) error

	// DeletePost deletes a post by his id
	DeletePost(User, PostId) error

	// Uploadpost add a post on your post list
	UploadPost(Post) (int64, error)

	// BannedCheck control if an user is banned by anotherone
	BanCheck(a User, b User) (bool, error)

	// CheckUser control if an user exist
	CheckUser(User) (bool, error)

	// SearchUser searches all the users that match the given name
	SearchUser(searcher User, userToSearch User) ([]User, error)

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
			user_id VARCHAR(15) PRIMARY KEY,
			username VARCHAR(15)
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id VARCHAR(15),
			publication_time DATETIME,
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
			comment_id INTEGER AUTOINCREMENT,
			text TEXT,
			time_comment TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
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
