package main

import (
    "log"
    "net/http"
    "go-crud/internal/router"
    "go-crud/internal/model"
    "go-crud/config"
)

func main() {
    config.LoadConfig()
    model.ConnectDatabase()
    r := router.InitRouter()
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
