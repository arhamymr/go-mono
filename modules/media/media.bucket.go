package mod_media

import (
	"context"
	"go-mono/buckets"
	"log"

	"cloud.google.com/go/storage"
)

func InitFirebaseStorage() (*storage.BucketHandle, context.Context) {
	app := buckets.ConnectBucket("firebase")

	ctx := context.Background()
	client, err := app.Firebase.Storage(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.Bucket("assets")
	if err != nil {
		log.Fatalln(err)
	}
	return bucket, ctx
}
