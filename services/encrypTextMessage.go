package services

import (
	"github.com/solrac97gr/cryptoAPI/models"
	"github.com/solrac97gr/cryptoAPI/util"
)

func EncryptTextMessage(messageToEncrypt models.Message) models.ReturnedMsg {
	encryptedMessage := new(models.ReturnedMsg)
	key := GetKey()
	message := messageToEncrypt.GetText()
	encrypted := util.Encrypt(message, key)
	encryptionkey := util.Convert(util.Key32)

	encryptedMessage.Message.SetText(message)
	encryptedMessage.DbMessage.SetEncryptedText(encrypted)
	encryptedMessage.DbMessage.SetEncryptionKey(encryptionkey)

	return *encryptedMessage
}
