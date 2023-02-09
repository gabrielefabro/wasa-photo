package database

func (db *appdbimpl) GetUserProfile (i id) (*Profile, error) {
	var profile Profile
	query := "SELECT User, Bio, Posts, FollowingCount, FollowerCount FROM profile WHERE id = ?"
	err := db.c.QueryRow(query, id).Scan(&profile.User, &profile.Bio, &profile.Posts, &profile.FollowingCount, &profile.FollowerCount)
	if err == sql.ErrNoRows {
		return nil, ProfileNotFound
	} else if err != nil {
		return nil, err
	}
	return &profile, nil
}