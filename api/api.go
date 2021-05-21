package api

import (
	"log"

	"github.com/solrac97gr/cryptoAPI/routes"

	"github.com/gofiber/fiber/v2"
)

//Init : Fiber Api Initialization and port definition
func Init() {
	app := fiber.New()
	app.Post("/decrypt", routes.DecryptTextMessage)
	app.Post("/encrypt", routes.EncryptTextMessage)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
