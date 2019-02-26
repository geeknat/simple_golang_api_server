// handlers.go

// Hosts all handlers to be executed on different routes
package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"strings"
)

// CreateCustomer handler receives form input,
// inserts a new customer record to the database
// and returns the Customer ID
func (a *App) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("first_name") == "" {
		respondWithJSON(w, http.StatusOK, 0, "First name is missing")
		return
	}

	if r.FormValue("last_name") == "" {
		respondWithJSON(w, http.StatusOK, 0, "Last name is missing")
		return
	}

	if r.FormValue("email") == "" {
		respondWithJSON(w, http.StatusOK, 0, "Email is missing")
		return
	}

	if r.FormValue("phone") == "" {
		respondWithJSON(w, http.StatusOK, 0, "Phone is missing")
		return
	}

	if r.FormValue("address") == "" {
		respondWithJSON(w, http.StatusOK, 0, "Address is missing")
		return
	}

	newCustomer := customer{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
		Address:   r.FormValue("address")}

	if err := newCustomer.insertCustomer(a.DB); err != nil {
		respondWithJSON(w, http.StatusOK, 0, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, 1, "Customer created successfully")

}

// FetchCustomer handler fetches details for a given customer
// based on their ID
func (a *App) FetchCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID, err := strconv.ParseInt(vars["id"], 0, 64)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, 0, "Missing customer ID")
		return
	}

	newCustomer := customer{CustomerID: customerID}
	if customerErr := newCustomer.getCustomer(a.DB); customerErr != nil {
		respondWithJSON(w, http.StatusOK, 0, customerErr.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, 1, newCustomer)

}

// APIAuthenticateMiddleware authenticates a request before
// passing it to the next handler
func (a *App) APIAuthenticateMiddleware(nextHandler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// We'll read a sample bearer token from Authorization
		// and perform a simple authentication

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {

			bearerToken := strings.Split(authorizationHeader, " ")

			if len(bearerToken) == 1 {

				if bearerToken[0] != "MY-SECRET-KEY" {
					respondWithJSON(w, http.StatusForbidden, 0, "Invalid authorization token")
					return
				}

				nextHandler(w, r)

			} else {
				respondWithJSON(w, http.StatusForbidden, 0, "Invalid authorization token")
				return
			}
		} else {
			respondWithJSON(w, http.StatusForbidden, 0, "Invalid authorization token")
			return
		}
	}
}
