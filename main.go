package main

import (
	"github.com/solrac97gr/cryptoAPI/api"
	"github.com/solrac97gr/cryptoAPI/database"
)

func main() {
	database.InitFirebase()
	api.Init()
}
