/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ProfileDoesNotExist = errors.New("profile does not exist")
var UserBanned = errors.New("user is banned")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// CreateUser create a new user
	CreateUser(User) error

	// SetMyUserName set a new Username for an existing profile
	SetMyUserName(User) error

	// GetUserProfile returns the profile matched with the ID
	GetUserProfile(id uint64) (Profile, error)

	// GetMyStream returns the stream of the id passed as argoument
	GetMyStream(id uint64) ([]Post, error)

	// GetMyFollowers returns the followers list
	GetMyFollowers(id uint64) ([]User, error)

	// GetMyFollowings returns the followings list
	GetMyFollowings(id uint64) ([]User, error)

	// GetMyBans returns the bans list
	GetMyBans(id uint64) ([]User, error)

	// GetLikes returns the likes list
	GetLikes(id uint64, postId uint64) ([]User, error)

	// GetComments returns the comments list
	GetComments(id uint64, postId uint64) ([]User, error)

	// FollowUser adds one profile from the followers list
	FollowUser(id uint64, secondId uint64) error

	// FollowUser removes one profile from the followers list
	UnfollowUser(id uint64, secondId uint64) error

	// BanUser adds one profile from the bans list
	BanUser(id uint64, secondId uint64) error

	// UnbanUser remove one profile from the bans list
	UnbanUser(id uint64, secondId uint64) error

	// LikePost add a like to the likes list
	LikePhoto(id uint64, postId uint64, secondId uint64) error

	// UnlikePost removes a like to the Unlikes list
	UnlikePhoto(id uint64, postId uint64, secondId uint64) error

	// CommentPhoto adds a comment in the comments list
	CommentPhoto(id uint64, postId uint64, comment string) error

	// UncommentPhoto adds a comment in the comments list
	UncommentPhoto(id uint64, postId uint64, commentId uint64) error

	// GetPost returns a post by his id
	GetPost(id uint64, postId uint64) (Post, error)

	// DeletePhoto deletes a post by his id
	DeletePhoto(id uint64, postId uint64) error

	// Uploadphoto add a post on your post list
	Uploadphoto(id uint64, img string, caption string) error

	// BannedCheck control if an user is banned by anotherone
	BanCheck(a User, b User) (bool, error)

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
	tables := [7]string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			username TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS profiles (
			user_id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(16),
			bio VARCHAR(64),
			following_count INTEGER,
			follower_count INTEGER,
			FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS posts (
			post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			publication_time DATETIME,
			bio VARCHAAR(64),
			like_count INTEGER,
			comment_count INTEGER,
			FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS  likes (
			post_id INTEGER NOT NULL,
			user_id VARCHAR(16) NOT NULL,
			PRIMARY KEY (id_photo,id_user),
			FOREIGN KEY(post_id) REFERENCES photos (post_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS comments (
			comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			post_id INTEGER
			text TEXT,
			time_comment TIMESTAMP,
			FOREIGN KEY (user) REFERENCES Users(user_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS banned_users (
			banner TEXT NOT NULL,
			banned TEXT NOT NULL,
			PRIMARY KEY (banner,banned),
			FOREIGN KEY(banner) REFERENCES users (user_id) ON DELETE CASCADE,
			FOREIGN KEY(banned) REFERENCES users (user_id) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS followers(
			follower TEXT NOT NULL,
			followed TEXT NOT NULL,
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
