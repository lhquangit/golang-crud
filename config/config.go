package config

import (
	"log"
	"github.com/spf13/viper"
)

var DB_url string
var JWTSecret string

func LoadConfig(){
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	DB_url = viper.GetString("DATABASE_URL")
	JWTSecret = viper.GetString("JWT_SECRET_KEY")
}
