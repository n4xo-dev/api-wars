package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iLopezosa/api-wars/rest/src/controllers"
)

func UsersRouter(api *fiber.Router) {
	u := (*api).Group("/users")
	u.Get("/", controllers.UserList)
	u.Get("/:id", controllers.UserRead) // Can use email as id
	u.Post("/", controllers.UserCreate)
	u.Delete("/:id", controllers.UserDelete)
	u.Put("/:id", controllers.UserPatch)
	u.Get("/:id/posts", controllers.UserPosts)
	u.Get("/:id/comments", controllers.UserComments)
	u.Get("/:id/messages", controllers.UserMessages)
	u.Get("/:userId/chat/:chatId/messages", controllers.UserChatMessages)
}
