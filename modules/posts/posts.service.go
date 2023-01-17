package mod_posts

import (
	"go-mono/configs"
)

func PostService() string {
	return "service posts"
}

func (data ModelPost) CreatePost() error {
	db := configs.ConnectDB("mysql")
	err := db.Create(&data.data).Error

	return err
}
