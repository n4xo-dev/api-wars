package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/graphql/graph"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func wrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
	}
}

func main() {
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

	// Start the server
	app.Listen(":3000")
}
