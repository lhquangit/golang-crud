package controller

import (
    // "encoding/json"
    "net/http"
    "go-crud/internal/repository"
    "go-crud/internal/handler"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := repository.GetAllUsers()
    if err != nil {
        handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Error getting users"})
        return
    }
    handler.ResponseWithJson(w, http.StatusOK, users)
}