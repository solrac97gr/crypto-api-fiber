package database

import (
	"fmt"
	"log"
)

func GetEncryptMessage(id string) {
	ref := DatabaseClient.NewRef("/text" + id)
	var data map[string]interface{}
	if err := ref.Get(FirebaseCtx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)
}
