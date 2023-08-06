package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/controllers"
)

func ChatsRouter(api *fiber.Router) {
	c := (*api).Group("/chats")
	c.Get("/", controllers.ChatList)
	c.Get("/:id", controllers.ChatRead) // Can load messages and participants w/ eager loading
	c.Post("/", controllers.ChatCreate)
	c.Post("/:id/users", controllers.ChatAddUsers)
	c.Delete("/:id", controllers.ChatDelete)
	c.Get("/:id/messages", controllers.ChatMessages)
	c.Get("/:chatId/user/:userId/messages", controllers.ChatUserMessages)
}
