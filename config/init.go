package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("config init...")

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file", err)
	}

	// app
	if err := app(); err != nil {
		log.Fatalln(err)
	}

	// redis
	if err := redis(); err != nil {
		log.Fatalln(err)
	}

	// mysql
	if err := mysql(); err != nil {
		log.Fatalln(err)
	}

	log.Println("config finished")
}
