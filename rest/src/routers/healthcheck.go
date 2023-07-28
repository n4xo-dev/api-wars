package routers

import "github.com/gofiber/fiber/v2"

func HealthcheckRouter(api *fiber.Router) {
	h := (*api).Group("/healthcheck")
	h.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "OK",
		})
	})
}
