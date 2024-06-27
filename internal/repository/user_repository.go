package repository

import (
	"go-crud/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := model.DB.Preload("Posts").Find(&users) // Preload các Post của User
	return users, result.Error
}

// Create new user
func (repo *UserRepository) CreateUser(user *model.User) error {
	return repo.DB.Create(user).Error
}