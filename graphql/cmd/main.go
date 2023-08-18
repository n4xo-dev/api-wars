package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/graphql/db"
	"github.com/iLopezosa/api-wars/graphql/graph"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func wrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
	}
}

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
	resetDB := flag.Bool("reset", false, "Reset the database")
	server := flag.Bool("server", false, "Run the server")
	flag.Parse()

	// Initialize the database
	fmt.Println("\nCreating connection to the databases...")

	db.Connect()
	db.RedisConnect()
	defer db.Disconnect()
	defer db.RedisDisconnect()

	if *resetDB {
		fmt.Println("\nResetting the database...")
		db.Reset()
		fmt.Println("\nSeeding database...")
		db.Seed()
	}

	// Initialize Fiber server
	if *server {
		app := fiber.New()

		// Create a gqlgen handler
		h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

		// Serve GraphQL API
		app.Post("/graphql", func(c *fiber.Ctx) error {
			wrapHandler(h.ServeHTTP)(c)
			return nil
		})

		// Serve GraphQL Playground
		app.Get("/", func(c *fiber.Ctx) error {
			wrapHandler(playground.Handler("GraphQL", "/graphql"))(c)
			return nil
		})

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
