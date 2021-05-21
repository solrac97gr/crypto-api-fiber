package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/cryptoAPI/models"
	"github.com/solrac97gr/cryptoAPI/services"
)

// DecryptTextMessage : Decrypt the information send it throw the api
func DecryptTextMessage(c *fiber.Ctx) error {
	encryptedMessage := new(models.Message)
	if err := c.BodyParser(encryptedMessage); err != nil {
		return err
	}
	decryptMessage := services.DecryptTextMessage(*encryptedMessage)
	return c.JSON(decryptMessage)
}
