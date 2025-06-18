package storage

import (
	"context"
	"log"

	"data-ingestion/model"

	"cloud.google.com/go/firestore"
)

type FirestoreClient interface {
	Collection(string) *firestore.CollectionRef
	Close() error
}

func StorePostsWithClient(ctx context.Context, client FirestoreClient, posts []model.Post) error {
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
