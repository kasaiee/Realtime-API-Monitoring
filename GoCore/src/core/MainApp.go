package main

import "fmt"


func main() {
	fmt.Println("Go Core Initialize")

	initPostgres()
	fmt.Printf(DB_NAME)
}