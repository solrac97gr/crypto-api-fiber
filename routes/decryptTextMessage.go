package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/cryptoAPI/models"
	"github.com/solrac97gr/cryptoAPI/services"
)

// DecryptTextMessage : Decrypt the information send it throw the api
func DecryptTextMessage(c *fiber.Ctx) error {
	ID := c.Params("id")
	encryptedMessage := new(models.ReturnedMsg)
	encryptedMessage.SetID(ID)
	error, decryptMessage := services.DecryptTextMessage(*encryptedMessage)
	if error {
		c.SendStatus(fiber.StatusBadRequest)
		return c.Send([]byte("The ID didn't Exist"))
	}
	return c.JSON(decryptMessage)
}
