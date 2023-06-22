package main

import (
	"fmt"
)

func main() {
	fmt.Println("Creating connection to the database...")

	db := dbConnect()

	fmt.Println("Resetting the database...")

	resetDB(db)
}
