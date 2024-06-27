package controller

import (
	"encoding/json"
	"go-crud/internal/dto"
	"go-crud/internal/handler"
	"go-crud/internal/model"
	"go-crud/internal/repository"
	"net/http"

	"gorm.io/gorm"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers()
	if err != nil {
		handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Error getting users"})
		return
	}
	handler.ResponseWithJson(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			handler.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
		}

		user := model.User{
			Username: req.Username,
			Role:     req.Role,
		}
		// Encode password
		if err := user.HashPassword(req.Password); err != nil {
			handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Failed to hash password"})
			return
		}

		userRepo := repository.UserRepository{DB: db}
		// Save user to DB
		if err := userRepo.CreateUser(&user); err != nil {
			handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Error creating user"})
			return
		}
		handler.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "User created successfully"})
	}
}
