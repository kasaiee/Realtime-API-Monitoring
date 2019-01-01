package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

import "fmt"

var DB_NAME, DB_USER, DB_PASS, DB_PORT string

func initPostgres() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_PORT = os.Getenv("DB_PORT")
	fmt.Printf(DB_NAME)
	fmt.Printf(DB_USER)
	fmt.Printf(DB_PASS)
	fmt.Printf(DB_PORT)
}
