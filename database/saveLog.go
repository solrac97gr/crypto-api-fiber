package database

import (
	"log"

	"github.com/solrac97gr/cryptoAPI/models"
)

func SaveLog(method string, url string) {
	ref := DatabaseClient.NewRef("/")
	logRef := ref.Child("log")
	newLog, err := logRef.Push(FirebaseCtx, nil)
	if err != nil {
		log.Fatalln("Error pushing child node:", err)
	}
	if err := newLog.Set(FirebaseCtx, &models.Log{
		Method: method,
		URL:    url,
	}); err != nil {
		log.Fatalln("Error setting value:", err)
	}
}
