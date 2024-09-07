package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n4xo-dev/api-wars/rest/controllers"
)

func MessagesRouter(api *fiber.Router) {
	m := (*api).Group("/messages")
	m.Get("/", controllers.MessageList)
	m.Get("/:id", controllers.MessageRead)
	m.Post("/", controllers.MessageCreate)
	m.Patch("/:id", controllers.MessagePatch)
	m.Put("/:id", controllers.MessageUpdate)
	m.Delete("/:id", controllers.MessageDelete)
}
