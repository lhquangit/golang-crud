package repository

import "go-crud/internal/model"

func GetAllPosts() ([]model.Post, error) {
	var post []model.Post
	result := model.DB.Preload("Tags").Find(&post)
	return post, result.Error
}

func CreatePost(post model.Post) error {
	result := model.DB.Create(&post)
	return result.Error
}