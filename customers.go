// customers.go

// Holds customer object and it's methods
// Database used for operation is MySQL

package main

import (
	"database/sql"
	"fmt"
)

type customer struct {
	CustomerID int64  `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

var customerTableFields = "customer_id,first_name,last_name,email,phone,address"

func (u *customer) getCustomer(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT "+customerTableFields+" FROM customers WHERE customer_id=%d", u.CustomerID)
	return db.QueryRow(statement).Scan(&u.CustomerID, &u.FirstName, &u.LastName, &u.Email, &u.Phone, &u.Address)
}

func (u *customer) insertCustomer(db *sql.DB) error {

	_, err := db.Exec("INSERT INTO customers (first_name,last_name,email,phone,address) "+
		"VALUES (?,?,?,?,?)", u.FirstName, u.LastName, u.Email, u.Phone, u.Address)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.CustomerID)

	if err != nil {
		return err
	}

	return nil
}
