/*
Package main is the entry point of the application. It initializes the database and runs the tests passed as command line
arguments.

Usage:

	$ go run src/cmd/main.go -tests=users,posts,comments,messages,chats

The flags are:

	-tests: comma-separated list of tests to run. The possible values are:
		- users
		- posts
		- comments
		- messages
		- chats

If no tests are passed, the application will only initialize the database.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/routers"
	"github.com/iLopezosa/api-wars/rest/src/tests"
	"github.com/joho/godotenv"
)

// Main function
func main() {
	// Initialize the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize command line arguments and flags
	testsPtr := flag.String("tests", "", "Tests to run")
	resetDB := flag.Bool("reset", false, "Reset the database")
	flag.Parse()

	// Initialize the database
	fmt.Println("\nCreating connection to the database...")

	db.Connect()
	defer db.Disconnect()

	if *resetDB {
		fmt.Println("\nResetting the database...")
		db.Reset()
		fmt.Println("\nSeeding database...")
		db.Seed()
	}

	// Run the tests
	if *testsPtr != "" {
		validateFlags(testsPtr)
		tests := strings.Split(*testsPtr, ",")
		runTests(tests)
	} else {
		fmt.Println("\nNo tests to run")
	}

	// Initialize Fiber server
	app := fiber.New()

	routers.Setup(app)

	app.Listen(":3000")

	fmt.Println("\nClosing connection to the database...")
}

// Validate the flags passed to the command line
func validateFlags(testsPtr *string) {
	re, err := regexp.Compile(`^(users|posts|comments|messages|chats)(,(users|posts|comments|messages|chats))*$`)
	if err != nil {
		log.Fatal(err)
	}
	if !re.MatchString(*testsPtr) {
		log.Fatal("Invalid tests flag")
	}
}

// runTests runs the tests passed as arguments
func runTests(toDo []string) {
	for _, t := range toDo {
		switch t {
		case "users":
			tests.TestUsersComplete()
		case "posts":
			tests.TestPostsComplete()
		case "comments":
			tests.TestCommentsComplete()
		case "messages":
			tests.TestMessagesComplete()
		case "chats":
			tests.TestChatsComplete()
		default:
			fmt.Printf("\nUnknown test: %s\n", t)
		}
	}
}
