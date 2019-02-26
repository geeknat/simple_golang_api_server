// app.go
package main

import (
	"fmt"
	"database/sql"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// App struct to hold database connection
type App struct {
	DB     *sql.DB
	Router *mux.Router
}

// Initialize prepares the server
func (a *App) Initialize() {
	connectionString := fmt.Sprintf(
		"%s:%s@/%s?parseTime=true",
		getEnvironmentVariable("DB_USER"),
		getEnvironmentVariable("DB_PASSWORD"),
		getEnvironmentVariable("DB_NAME"))

	var err error

	// Open connection to the database
	a.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run starts the server at given address
func (a *App) Run(addr string) {
	if err := http.ListenAndServe(addr, a.Router); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port " + addr)
}

// initializeRoutes lists all the routes for the API
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/customer", a.APIAuthenticateMiddleware(a.CreateCustomer)).Methods("POST")
	a.Router.HandleFunc("/customer/{id:[0-9]+}", a.FetchCustomer).Methods("GET")
}
