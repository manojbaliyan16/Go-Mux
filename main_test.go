package main_test

import (
	"fmt"
	"github.com/manojbaliyan16/Go-Mux"
	"log"
	"os"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	fmt.Println("In testMain ")
	a.Initialize(
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_DATABASE"))
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	fmt.Println("In-->ensureTableExists ")
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	fmt.Println("In clearTable ")
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`
