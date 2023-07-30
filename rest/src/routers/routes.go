package routers

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	api := app.Group("/rest")

	UsersRouter(&api)
	PostsRouter(&api)
	CommentsRouter(&api)
	MessagesRouter(&api)
	ChatsRouter(&api)
	HealthcheckRouter(&api)
}
