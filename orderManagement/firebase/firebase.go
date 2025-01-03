package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/option"
	secretspb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var Client *firestore.Client

func InitFirebase() {
	ctx := context.Background()

	// Accede al secreto desde Google Secret Manager
	credentials, err := getFirebaseCredentials(ctx, "projects/task-444104/secrets/FIREBASE_CREDENTIALS/versions/latest")
	if err != nil {
		log.Fatalf("Failed to get Firebase credentials from Secret Manager: %v", err)
	}

	// Inicializa el cliente de Firestore con las credenciales obtenidas
	Client, err = firestore.NewClient(ctx, "task-444104", option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
}

// Obtiene las credenciales de Firebase desde Secret Manager
func getFirebaseCredentials(ctx context.Context, secretName string) ([]byte, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Accede al secreto
	req := &secretspb.AccessSecretVersionRequest{
		Name: secretName,
	}

	resp, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Data, nil
}

// func InitFirebaseLocal() {
// 	ctx := context.Background()
// 	// Carga las credenciales desde el archivo key.json
// 	opt := option.WithCredentialsFile("key.json")
// 	client, err := firestore.NewClient(ctx, "task-444104", opt)
// 	if err != nil {
// 		log.Fatalf("Failed to create Firestore client: %v", err)
// 	}
// 	Client = client
// }
