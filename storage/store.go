package storage

import (
	"context"
	"log"
	"os"

	"data-ingestion/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func StorePosts(ctx context.Context, posts []model.Post) error {
	projectID := os.Getenv("GCP_PROJECT_ID")
	credsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credsPath))
	if err != nil {
		return err
	}
	defer client.Close()

	for _, post := range posts {
		_, _, err := client.Collection("posts").Add(ctx, post)
		if err != nil {
			log.Printf("Failed to store post ID %d: %v", post.Id, err)
		} else {
			log.Printf("Successfully stored post ID %d", post.Id)
		}
	}
	return nil
}
