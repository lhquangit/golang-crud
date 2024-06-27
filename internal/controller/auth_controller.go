package controller

import (
	"encoding/json"
	"go-crud/auth"
	"go-crud/internal/dto"
	"go-crud/internal/handler"
	"go-crud/internal/model"
	"net/http"

	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    var req dto.LoginRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        handler.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
        return
    }

    var user model.User
    if err := db.Where("username=?", req.Username).First(&user).Error; err != nil {
        handler.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid username"})
        return
    }

    if !user.CheckPassword(req.Password) {
        handler.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid password"})
        return
    }
    // Parse and validate request body
    // Authenticate user
    // Generate JWT token
    token, err := auth.GenerateJWT(user.Username, user.Role)
    if err != nil {
        handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
        return
    }
    handler.ResponseWithJson(w, http.StatusOK, map[string]string{"token": token})
}
