package main

import (
	"github.com/solrac97gr/cryptoAPI/api"
	cacheDb "github.com/solrac97gr/cryptoAPI/cache"
	"github.com/solrac97gr/cryptoAPI/database"
)

func main() {
	cacheDb.InitCache()
	database.InitFirebase()
	api.Init()
}
