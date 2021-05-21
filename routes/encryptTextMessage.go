package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/cryptoAPI/models"
	"github.com/solrac97gr/cryptoAPI/services"
)

// EncryptTextMessage : Encrypt the information send it throw the api
func EncryptTextMessage(c *fiber.Ctx) error {
	message := new(models.Message)
	if err := c.BodyParser(message); err != nil {
		return err
	}
	valErr, errMsg := message.Validate()
	if valErr {
		c.Send([]byte(errMsg))
	}
	encryptedMessage := services.EncryptTextMessage(*message)
	return c.JSON(encryptedMessage)
}
