package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongoURL string
)

func LoadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoURL = os.Getenv("MONGO_URL") // Get url from env
}
