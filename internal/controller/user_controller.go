package controller

import (
	"encoding/json"
	"go-crud/internal/handler"
	"go-crud/internal/model"
	"go-crud/internal/repository"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := repository.GetAllUsers()
    if err != nil {
        handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Error getting users"})
        return
    }
    handler.ResponseWithJson(w, http.StatusOK, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var newUser model.User

    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        handler.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
        return
    }

    if err := repository.CreateUser(newUser); err != nil {
        handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Error creating user"})
        return
    }

    handler.ResponseWithJson(w, http.StatusCreated, newUser)
}