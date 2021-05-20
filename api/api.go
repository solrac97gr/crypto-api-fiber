package api

import (
	"github.com/gofiber/fiber/v2"
)

//Api init
func Init() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}
