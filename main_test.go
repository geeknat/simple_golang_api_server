// handlers_test.go
package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"os"
	"log"
	"strconv"
	"encoding/json"
)

var a App

const tableCustomerCreationQuery = `CREATE TABLE IF NOT EXISTS customers
(
    customer_id INT AUTO_INCREMENT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
 	phone TEXT NOT NULL,
 	email TEXT NOT NULL,
 	address TEXT NOT NULL
)`

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()

	ensureTableExists()

	code := m.Run()

	os.Exit(code)
}

// TestCreateCustomer performs tests on CreateCustomer handler
func TestCreateCustomer(t *testing.T) {
	data := url.Values{}
	data.Set("first_name", "Dennis")
	data.Set("last_name", "Natalia")
	data.Set("email", "geeknat7@gmail.com")
	data.Set("phone", "+254718353279")
	data.Set("address", "Nairobi Kenya")

	req, err := http.NewRequest("POST", "/customer", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "MY-SECRET-KEY")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	rw := httptest.NewRecorder()
	a.Router.ServeHTTP(rw, req)

	// Check the status code is what we expect.
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var m response
	json.Unmarshal(rw.Body.Bytes(), &m)

	if m.Success != 1 {
		t.Errorf("Expected success to be 1, Got %d with message '%s'",
			m.Success, m.Message)
	}

}

// TestFetchCustomer performs tests on FetchCustomer handler
func TestFetchCustomer(t *testing.T) {

	req, err := http.NewRequest("GET", "/customer/5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	a.Router.ServeHTTP(rw, req)

	// Check the status code is what we expect.
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var m response
	json.Unmarshal(rw.Body.Bytes(), &m)

	if m.Success != 1 {
		t.Errorf("Expected success to be 1, Got %d with message '%s'",
			m.Success, m.Message)
	}

	t.Logf("Message %v", m.Message)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCustomerCreationQuery); err != nil {
		log.Fatal(err)
	}
}
