package database

import (
	"fmt"
	"os"
	"strconv"

	"git.gabrielefabro.it/service/api/utils"
)

// Database function that removes a photo from the database
func (db *appdbimpl) DeletePost(user User, postId PostId) error {

	_, err := db.c.Exec("DELETE FROM posts WHERE user_id = ? AND post_id = ? ",
		user.User_id, postId.Post_id)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	user_id_i, err := strconv.Atoi(user.User_id)
	if err != nil {
		fmt.Println("Errore durante la conversione:", err)
		return err
	}

	var post_id = postId.Post_id

	// Delete file
	err = os.Remove(utils.GetPostPhotoPath(user_id_i, int64(post_id)))
	if err != nil {
		return err
	}

	return nil
}
