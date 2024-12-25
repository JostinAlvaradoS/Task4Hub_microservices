package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirebase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("key.json")
	var err error
	Client, err = firestore.NewClient(ctx, "task-444104", sa)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
}
