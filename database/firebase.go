package database

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	configCred "github.com/solrac97gr/cryptoAPI/config"
	"google.golang.org/api/option"
)

var App *firebase.App
var DatabaseClient *db.Client
var FirebaseCtx context.Context

func InitFirebase() {
	ctx := context.Background()
	FirebaseCtx = ctx
	opt := option.WithCredentialsJSON([]byte(configCred.GetFirebaseCredentials()))
	conf := &firebase.Config{
		DatabaseURL: configCred.GetFirebaseDatabaseUrl(),
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	App = app
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	DatabaseClient = client
}
