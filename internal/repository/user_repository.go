package repository

import (
	"go-crud/internal/model"
)

func GetAllUsers() ([]model.User, error) {
	var users []model.User
    result := model.DB.Preload("Posts").Find(&users) // Preload các Post của User
    return users, result.Error
}