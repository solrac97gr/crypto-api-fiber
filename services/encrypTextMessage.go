package services

import (
	"github.com/solrac97gr/cryptoAPI/database"
	"github.com/solrac97gr/cryptoAPI/models"
	"github.com/solrac97gr/cryptoAPI/util"
)

func EncryptTextMessage(messageToEncrypt models.Message) models.ReturnedMsg {
	encryptedMessage := new(models.ReturnedMsg)
	key := GetKey()
	message := messageToEncrypt.GetText()
	encrypted := util.Encrypt(message, key)
	encryptionkey := util.ConvertToByteString(util.Key32)
	ID := database.SaveEncryptMessage(key, encrypted)

	encryptedMessage.SetID(ID)
	encryptedMessage.SetText(message)
	encryptedMessage.SetEncryptedText(encrypted)
	encryptedMessage.SetEncryptionKey(encryptionkey)

	return *encryptedMessage
}
