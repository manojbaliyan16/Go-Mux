package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	fmt.Println("In Initialize ")
	//fmt.Printf("%s", connectionString)
	var err error
	a.DB, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		fmt.Println(" ")
		log.Fatal(err.Error())
	}
	a.Router = mux.NewRouter()
}
func (a *App) Run(addr string) {
	fmt.Println("Running the app ")
}
