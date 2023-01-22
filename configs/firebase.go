package configs

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func FirebaseConfig() *firebase.App {
	const serviceKey = "credentials/firebase.json"
	config := &firebase.Config{
		StorageBucket: GET("FIREBASE_STORAGE_BUCKET"),
	}

	opt := option.WithCredentialsFile(serviceKey)
	app, err := firebase.NewApp(context.Background(), config, opt)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return app
}
