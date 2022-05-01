package main

import (
	"context"
	"fmt"
	"log"
)

var mg MongoInstance

func main() {

	// Load var from .env file
	LoadVar()

	err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Disconnect")
		err := mg.Client.Disconnect(context.TODO())
		if err != nil {
			return
		}
	}()

	// Create the app
	app := New()
	// Listen to port 1812
	log.Fatal(app.Listen(":1815"))
}
