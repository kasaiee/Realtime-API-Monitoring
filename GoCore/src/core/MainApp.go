package main

import "fmt"


func main() {
	fmt.Println("Go Core Initialize")

	initPostgres()
	fmt.Printf(DB_NAME)
	fmt.Printf(DB_USER)
	fmt.Printf(DB_PASS)
	fmt.Printf(DB_PORT)
}