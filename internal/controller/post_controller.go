package controller

import (
	"go-crud/internal/handler"
	"go-crud/internal/repository"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, error := repository.GetAllPosts()

	if error != nil {
		handler.ResponseWithJson(w, http.StatusInternalServerError, map[string]string{"message": "Error getting users"})
		return
	}

	handler.ResponseWithJson(w, http.StatusOK, posts)
}
