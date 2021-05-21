package database

import (
	"log"

	"github.com/solrac97gr/cryptoAPI/models"
)

func SaveEncryptMessage(encryptionKey string, encryptedText string) {
	ref := DatabaseClient.NewRef("/text")
	err := ref.Set(FirebaseCtx, &models.DbMessage{
		EncryptedText: encryptedText,
		EncryptionKey: encryptionKey,
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}

}
