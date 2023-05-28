package utils

import (
	"fmt"
)

func GetPostPhotoPath(user_id int, post_id int) string {
	return fmt.Sprintf("./storage/%d/posts/%d.jpeg", user_id, post_id)
}