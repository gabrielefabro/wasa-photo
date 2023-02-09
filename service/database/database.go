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
	"time"
)

// Profile struct represent a profile.
type Profile struct {
	User           uint64
	Bio            string
	Posts          uint64
	PostCount      uint64
	FollowingCount uint64
	FollowerCount  uint64
}

// Post struct represent a post.
type Post struct {
	User            uint64
	PublicationTime time.Time
	Bio             string
	LikeCount       uint64
	CommentCount    uint64
	PhotoUrl        string
}

// Comment struct represent a comment
type Comment struct {
	User        uint64
	Text        string
	CommentId   uint64
	TimeComment time.Time
}

// CommentList represent a list of profile
type Comments struct {
	Comments []Comment
}

// ProfileList represent a list of profile
type Profiles struct {
	Profiles []Profile
}

// ProfileList represent a list of profile
type Posts struct {
	Posts []Post
}

// User represent the couple ID and UserName
type User struct {
	ID       uint64
	UserName string
}

// ProfileList represent a list of profile
type Users struct {
	users []User
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// SetMyUserName set a new Username for an existing profile
	SetMyUserName(User) error

	// GetUserProfile returns the profile matched with the ID
	GetUserProfile(id uint64) (Profile, error)

	// GetMyStream returns the stream of the id passed as argoument
	GetMyStream(id uint64) ([]Posts, error)

	// GetMyFollowers returns the followers list
	GetMyFollowers(id uint64) ([]Users, error)

	// GetMyFollowings returns the followings list
	GetMyFollowings(id uint64) ([]Users, error)

	// GetMyBans returns the bans list
	GetMyBans(id uint64) ([]Users, error)

	// GetLikes returns the likes list
	GetLikes(id uint64, postId uint64) ([]Users, error)

	// GetComments returns the comments list
	GetComments(id uint64, postId uint64) ([]Users, error)

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

	// GetPost returns a post by his id
	GetPost(id uint64, postId uint64) (Post, error)

	// DeletePhoto deletes a post by his id
	DeletePhoto(id uint64, postId uint64) error

	// Uploadphoto add a post on your post list
	Uploadphoto(id uint64, img string, caption string) error

	// CommentPhoto adds a comment in the comments list
	CommentPhoto(id uint64, postId uint64, comment string) error

	// UncommentPhoto adds a comment in the comments list
	UncommentPhoto(id uint64, postId uint64, commentId uint64) error

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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
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
