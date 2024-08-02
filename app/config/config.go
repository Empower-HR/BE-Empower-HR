package config

import (
	"log"
	"os"
	"strconv"
)

var (
	JWT_SECRET     string
	CLOUDINARY_URL string
	EMAIL_PASSWORD string
	EMAIL_SENDER   string
	API_KEYS       string
)

type AppConfig struct {
	USER     string
	PASSWORD string
	HOST     string
	PORT     int
	DBNAME   string
}

func ReadEnv() *AppConfig {
	var app = AppConfig{}
	app.USER = os.Getenv("DBUSER")
	app.PASSWORD = os.Getenv("DBPASS")
	app.HOST = os.Getenv("DBHOST")
	portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
	if errConv != nil {
		// panic("error convert dbport")
		log.Println("Err", errConv.Error())
	}
	app.PORT = portConv
	app.DBNAME = os.Getenv("DBNAME")
	JWT_SECRET = os.Getenv("JWTSECRET")
	CLOUDINARY_URL = os.Getenv("CLOUDINARY_URL")
	EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
	EMAIL_SENDER = os.Getenv("EMAIL_SENDER")
	API_KEYS = os.Getenv("API_KEYS")
	return &app
}

func InitConfig() *AppConfig {
	return ReadEnv()
}
