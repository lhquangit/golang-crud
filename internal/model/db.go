package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"go-crud/config"
)

var DB *gorm.DB
var models = []interface{} {
	&User{},
	&Post{},
	&Tag{},
}

func ConnectDatabase(){
	var err error
	DB, err = gorm.Open(
		postgres.Open(config.DB_url),
		&gorm.Config{},
	)
	if err != nil{
		log.Fatal("Failed to connect to database:", err)
		return
	}
	err = DB.AutoMigrate(models...) //auto migrate tables in database follow the defined structs
	if err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
	}
}