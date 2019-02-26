// main.go

// Simple API server that allows customer registration and
// fetching of customer details using customer ID
//
// Routing package used is gorilla/mux router
// Database used is MySQL
// Built with go1.11.5

package main

import (
	"log"
	"github.com/subosito/gotenv"
)

// Check if there exists a .env file
func init() {
	if err := gotenv.Load(); err != nil {
		log.Println("File .env not found")
		return
	}
}

func main() {

	app := App{}

	app.Initialize()

	app.Run(":83")
}
