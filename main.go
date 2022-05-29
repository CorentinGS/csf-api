package main

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var mg MongoInstance
var c *cache.Cache

func main() {

	// Load var from .env file
	LoadVar()

	c = cache.New(5*60*60*time.Second, 10*60*60*time.Second)

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
