package api

import (
	"log"

	middlewares "github.com/solrac97gr/cryptoAPI/middleware"
	"github.com/solrac97gr/cryptoAPI/routes"

	"github.com/gofiber/fiber/v2"
)

//Init : Fiber Api Initialization and port definition
func Init() {
	app := fiber.New()
	app.Use(middlewares.RequestLogger)
	v1 := app.Group("/v1")
	v1.Get("/decrypt/:id", routes.DecryptTextMessage)
	v1.Post("/encrypt", routes.EncryptTextMessage)

	err := app.Listen(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
