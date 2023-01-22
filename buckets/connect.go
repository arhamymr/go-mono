package buckets

import (
	"go-mono/configs"

	firebase "firebase.google.com/go/v4"
)

type Buckets struct {
	Firebase *firebase.App
}

func ConnectBucket(provider string) Buckets {
	bucket := Buckets{}

	switch provider {
	case "firebase":
		bucket.Firebase = configs.FirebaseConfig()
	default:
		panic("bucket not defined")
	}

	return bucket
}
