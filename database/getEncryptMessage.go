package database

import (
	"log"

	"github.com/solrac97gr/cryptoAPI/models"
)

func GetEncryptMessage(id string) (bool, models.ReturnedMsg) {
	ref := DatabaseClient.NewRef("/text/" + id)
	var response models.ReturnedMsg
	if err := ref.Get(FirebaseCtx, &response); err != nil {
		log.Fatalln("Error reading value:", err)
	}
	if len(response.GetEncryptedText()) <= 0 {
		return true, response
	}
	ID := ref.Key
	response.SetID(ID)
	return false, response
}
