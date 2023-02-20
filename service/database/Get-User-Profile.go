package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetUserProfile(id uint64) (Profile, error) {
	var ProfileNotFound = errors.New("Profile not found")
	var profile Profile
	query := "SELECT User, Bio, Posts, FollowingCount, FollowerCount FROM profile WHERE id = ?"
	err := db.c.QueryRow(query, id).Scan(&profile.User, &profile.Bio, &profile.Posts, &profile.FollowingCount, &profile.FollowerCount)
	if err == sql.ErrNoRows {
		return Profile{}, ProfileNotFound
	} else if err != nil {
		return Profile{}, err
	}
	return profile, nil
}
