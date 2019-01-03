package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// import postgres package
import (
	"github.com/go-pg/pg"
)

var DB_NAME,
	DB_USER,
	DB_PASS,
	DB_PORT string //TODO: may init locally

func initPostgres() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_PORT = os.Getenv("DB_PORT")
}

func connectPostgres() {
	db := pg.Connect(&pg.Options{
		User:     DB_USER,
		Password: DB_PASS,
		Database: DB_NAME,
	})
	defer db.Close()
}

func selectRequests(){
	fmt.Println("Select requests based on django project")
}
