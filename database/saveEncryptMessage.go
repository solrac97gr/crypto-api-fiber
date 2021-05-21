package database

import (
	"log"

	"github.com/solrac97gr/cryptoAPI/models"
)

func SaveEncryptMessage(encryptionKey string, encryptedText string) string {
	ref := DatabaseClient.NewRef("/")
	textRef := ref.Child("text")
	newText, err := textRef.Push(FirebaseCtx, nil)
	if err != nil {
		log.Fatalln("Error pushing child node:", err)
	}
	if err := newText.Set(FirebaseCtx, &models.DbMessage{
		EncryptedText: encryptedText,
		EncryptionKey: encryptionKey,
	}); err != nil {
		log.Fatalln("Error setting value:", err)
	}
	textID := newText.Key
	return textID
}
