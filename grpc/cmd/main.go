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
	"net"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/n4xo-dev/api-wars/grpc/services"
	"github.com/n4xo-dev/api-wars/grpc/tests"
	"github.com/n4xo-dev/api-wars/lib/db"
	"google.golang.org/grpc"
)

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
	fmt.Println("\nCreating connection to the databases...")

	db.Connect()
	defer db.Disconnect()
	db.RedisConnect()
	defer db.RedisDisconnect()

	// Reset and seed the database
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

	// Initialize gRPC server
	if *server {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		fmt.Println("\nStarting gRPC server on port 50051...")
		grpcServer := grpc.NewServer()
		services.RegisterServices(grpcServer)

		PrintMemUsage()

		go func() {
			grpcServer.Serve(lis)
		}()

		_ = <-sigs

		grpcServer.Stop()
	}

	fmt.Println("\nClosing connection to the database...")
}
