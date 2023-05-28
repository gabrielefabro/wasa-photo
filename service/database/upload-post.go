package database

import (
	"fmt"
	"os"
	"strconv"

	"git.gabrielefabro.it/service/api/utils"
)

// Function that creates a post on the database
func (db *appdbimpl) UploadPost(post Post, data []byte) (int64, error) {

	var query = "INSERT INTO posts (user_id, username, publication_time) VALUES (?,?,?)"

	res, err := db.c.Exec(query, post.User_id, post.Username, post.Publication_time)

	if err != nil {
		return -1, err
	}

	post_id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	user_id := post.User_id

	user_id_i, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Println("Errore durante la conversione:", err)
		return -1, err
	}

	path := utils.GetPostPhotoPath(user_id_i, post_id)

	// Save the image
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		return -1, err
	}

	// Crop the image
	err = utils.SaveAndCrop(path, 720, 720)
	if err != nil {
		return -1, err
	}

	return post_id, nil
}
