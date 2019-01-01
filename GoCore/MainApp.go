package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

import "fmt"


func main() {
	fmt.Println("Go Core Initialize")
	
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("DB_NAME")
	fmt.Printf(s3Bucket)
}