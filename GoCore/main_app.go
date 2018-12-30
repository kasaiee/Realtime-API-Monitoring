package main
import(
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
}