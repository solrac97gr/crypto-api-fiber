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
	opt := option.WithCredentialsJSON([]byte(`{
		"type": "service_account",
		"project_id": "gotest-42c91",
		"private_key_id": "be83b4305c299a6967f6dea9ee9f6068e1498d3c",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCzGidG5FMfc1rL\nymfRIcsEpJrHYOw/8Xnn09rj3ftKOz3y+7GayPusaoLnJGeW7MDax8Vm0s5Q4+vR\n4s8ABnl/zQn3nX1ieY5vx2WQ84dakqxqB/VXb6ku+VWeOtZaFXcnWFOuCIgi/aGA\nX4ts86YDMCAOPhD7+CCnWV7vATeE65Pic0YhtFNcmHwg3U56XYw/p6yEZjk9ESW4\n7DhZlD+0b0oJ/KJh4XghwjyR4RJpTNq6G/PuN5tC4wSZlY1AcJEYBJUeen+g4K/D\n9YNZPEZU6ebka6Te7Bh2zLKnDE3oLNVFGo5kmjEkUgsrHX7h3AB6Sc0+5A2f0CfS\nQ3DTqCiJAgMBAAECggEAJ5vBkup8vYhXzJjbI8Nq8ABtIGm9ZZjEInJWRq15XN3E\nIf/qs7X9/o/hmjeRVy3NqrSiYc5+BSKarOSRbcqAxTQ1n5arO4NxbiP1QLyysuOc\nEKjo56jQjAxlFGSjsfFsU/2zJ9pILSDVPhwjK/moPa0/cZblT6v8oLXhik7FFfZZ\nzAh0Yu3nMOthXNdiu7TI8Ep56nAdNvo5Uiv+jhSmx0FXgIMsENF0KywCPYAWoFvJ\nXspHmPXJxKyZJ8i1sIBW+LyEmuNJpAfOUMTxj05yvNtJA+cOybYAOHK8LsU5tWqC\nJy8v7mXdxU4jeo0eX0CyAL0ydywnmEImDiN/OOd9+QKBgQD0JcNkBLslUf/HfO1H\nbdQZWqvFQTsNSPGBA7tMEe8ZfaVoEAmYatXiKyZ5oTd3VYKpFL/HYah+99k5v3UG\n7h34xR46hbpRobG5hHyccAMy3iJ1X8bbz3jDgHxxbBxLoerINaJrc6NPd3EfZI/b\nDfvlaV0LuDpdi9jbkzzUWDhR1QKBgQC7zAOgGXQKzD6Z6Yaao9+6/m3e7UpIU+cd\nwTkSCs9OArc3lUaH7fUOh4Nn17Cf5nfGHtpccna9gTsg35ZbCxP61p9cIpGgyt3G\ntxmNAqPW/S+a3TbGu6AOWJU30NG+WTluIiRZfgH61VZm61gB4IVf74d7JBL9jMJ0\nAnMHnpCh5QKBgDxq4/FqYk2X/y4EsktM2IH1uDA2NkDRXkJJBB6NCZhLSZV48obv\nWUKA3yKWTL2RQaZJ5jOW7TWObAkoYfN3FZ+sWBCvay4sT2jTamrkJUZ/RhO5weI/\nc2kd8K24zT/CL/GfpDPnB/DwY+Qa5KaWzVE2Q7pUuszPgLjFtPXvZvCJAoGBALAv\nlhMk7h6XQgAJ/DOmnOcFfrTP8Ins1X6v4cgQhi8NEzE9hh9O20LXViwdSmKpySMB\nnSFwbgAadwu/WINv4D3MD1RURkYRLYIaFAmcOXQgswHammJWcrIi31OvQKgRc3Tk\n6S9graDTd601DKsuJlM8GOIplPt1jJVkGuFCHTUVAoGBAI+MGkeRsWBg9Or5IjQl\nyPI/dkWtNL2ktyA0y/lfIuo1U8HJ988FUcL0W/R11azTC3VGTUzrulocaFspXKhX\nEUXXLv8Yyxi/rbWvau7OaikllbGz53TyBy9WQVNF+DCjoq+SH58Q5RuS5vB0kOxZ\nYRn/k7NzKZs7O9z2jWCO80pz\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-gm41l@gotest-42c91.iam.gserviceaccount.com",
		"client_id": "109923863577138692632",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-gm41l%40gotest-42c91.iam.gserviceaccount.com"
	  }
	  `))
	conf := &firebase.Config{
		//Replace Database URL here
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
