package mod_posts

import (
	"go-mono/configs"
)

func PostService() string {
	return "service posts"
}

func (data ModelPost) CreatePost() error {
	err := configs.DB.Create(&data.data).Error

	return err
}
