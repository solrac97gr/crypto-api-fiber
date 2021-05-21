package database

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var App *firebase.App
var DatabaseClient *db.Client
var FirebaseCtx context.Context

func InitFirebase() {
	ctx := context.Background()
	FirebaseCtx = ctx
	//Full Path credentials
	opt := option.WithCredentialsFile("/Users/solrac97g/go/src/github/solrac97gr/cryptoAPI/serviceAccountKey.json")
	conf := &firebase.Config{
		DatabaseURL: "https://gotest-42c91-default-rtdb.firebaseio.com/",
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
