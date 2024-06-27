package router

import (
	"go-crud/internal/controller"
	"go-crud/internal/model"
	"go-crud/middleware"
	"net/http"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()
    // Định nghĩa route cho trang chủ để kiểm tra
    r.HandleFunc("/", HomeHandler).Methods("GET")

    // Public routes
    r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
                        controller.Login(w,r,model.DB)}).Methods("POST")
 
    // Protected routes
    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.JWTAuthMiddleware)
    
    // Admin routes
    admin := api.PathPrefix("/admin").Subrouter()
    admin.Use(middleware.AdminOnly)
    admin.HandleFunc("/create_user", controller.CreateUser(model.DB)).Methods("POST") // Create user
    
    // User routes
    // api.Use(middleware.Authenticate)
    api.HandleFunc("/users", controller.GetAllUsers).Methods("GET") // Get all users
    api.HandleFunc("/posts", controller.GetAllPosts).Methods("Get") // Get all posts
    return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the Home Page!"))
}