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
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/db"
	"github.com/iLopezosa/api-wars/rest/src/routers"
	"github.com/iLopezosa/api-wars/rest/src/tests"
	"github.com/joho/godotenv"
)

// Main function
func main() {
	defer PrintMemUsage()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	// Initialize the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize command line arguments and flags
	testsPtr := flag.String("tests", "", "Tests to run")
	resetDB := flag.Bool("reset", false, "Reset the database")
	server := flag.Bool("server", false, "Run the server")
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
		if err := validateFlags(testsPtr); err == nil {
			tests.Run(testsPtr)
		} else {
			fmt.Println("\nERROR:", err)
		}
	} else {
		fmt.Println("\nNo tests to run")
	}

	// Initialize Fiber server
	if *server {
		app := fiber.New()
		routers.Setup(app)
		PrintMemUsage()
		go func() {
			app.Listen(":3000")
		}()
		_ = <-sigs
		app.Shutdown()
	}

	fmt.Println("\nClosing connection to the database...")
}

// Validate the flags passed to the command line
func validateFlags(testsPtr *string) error {
	re, err := regexp.Compile(`(^(users|posts|comments|messages|chats)(,(users|posts|comments|messages|chats))*$)|(^all$)`)
	if err != nil {
		return err
	}
	if !re.MatchString(*testsPtr) {
		return errors.New("Invalid tests flag")
	}
	return nil
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the
// number of garage collection cycles completed. For info on each,
// see: https://golang.org/pkg/runtime/#MemStats
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
