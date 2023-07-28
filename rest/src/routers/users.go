package routers

import "github.com/gofiber/fiber/v2"

func UsersRouter(api *fiber.Router) {
	u := (*api).Group("/users")
	u.Get("/", UserList)
	u.Get("/:id", UserRead) // Can use email as id
	u.Post("/", UserCreate)
	u.Delete("/:id", UserDelete)
	u.Patch("/:id", UserPatch)
	u.Get("/:id/posts", UserPosts)
	u.Get("/:id/comments", UserComments)
	u.Get("/:id/messages", UserMessages)
	u.Get("/:id/chat/:id2/messages", UserChatMessages)
}
