package services

import (
	"github.com/solrac97gr/cryptoAPI/models"
	"github.com/solrac97gr/cryptoAPI/util"
)

func DecryptTextMessage(encryptedMessage models.Message) models.Message {
	decryptMessage := new(models.Message)
	key := GetKey()
	message := encryptedMessage.GetText()
	decryptMessage.SetText(util.Decrypt(key, message))
	return *decryptMessage
}
