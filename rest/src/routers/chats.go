package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/controllers"
)

func ChatsRouter(api *fiber.Router) {
	c := (*api).Group("/chats")
	c.Get("/", controllers.ChatList)
	c.Get("/:id", controllers.ChatRead) // Can load messages and participants w/ eager loading
	c.Post("/", controllers.ChatCreate)
	c.Patch("/:id", controllers.ChatUpdate)
	c.Delete("/:id", controllers.ChatDelete)
	c.Get("/:id/messages", controllers.ChatMessages)
	c.Get("/:id/user/:id2/messages", controllers.ChatUserMessages)
}