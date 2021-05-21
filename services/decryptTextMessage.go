package services

import (
	"github.com/solrac97gr/cryptoAPI/database"
	"github.com/solrac97gr/cryptoAPI/models"
	"github.com/solrac97gr/cryptoAPI/util"
)

func DecryptTextMessage(encryptedMessage models.ReturnedMsg) (bool, models.ReturnedMsg) {
	error, encryptedMsgFromDB := database.GetEncryptMessage(encryptedMessage.GetID())
	if error {
		return true, encryptedMsgFromDB
	}
	key := encryptedMsgFromDB.GetEncryptionKey()
	message := encryptedMsgFromDB.GetEncryptedText()
	decryptedMessage := util.Decrypt(message, key)
	encryptionKey := util.ConvertToByteString(util.Key32)

	encryptedMsgFromDB.SetText(decryptedMessage)
	encryptedMsgFromDB.SetEncryptionKey(encryptionKey)
	return false, encryptedMsgFromDB
}
