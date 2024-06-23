package router

import (
    "github.com/gorilla/mux"
    "go-crud/internal/controller"
    "net/http"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()
    // Định nghĩa route cho trang chủ để kiểm tra
    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/api/users", controller.GetAllUsers).Methods("GET")
    // r.HandleFunc("/api/users", controller.CreateUser).Methods("POST")
    // Các route khác cho Post và Tag
    return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the Home Page!"))
}