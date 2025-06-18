package main

import (
	"context"
	"log"
	"time"

	collector "data-ingestion/collection"
	storage "data-ingestion/storage"
	transformer "data-ingestion/transformation"
)

func main() {
	apiURL := "https://jsonplaceholder.typicode.com/posts"
	timeout := 10 * time.Second
	posts, err := collector.GetPosts(apiURL, timeout)
	if err != nil {
		log.Fatalf("Failed to get posts: %v", err)
	}
	transformed := transformer.TransformPosts(posts, "placeholder_api")
	ctx := context.Background()
	if err := storage.StorePosts(ctx, transformed); err != nil {
		log.Fatalf("Failed to store posts: %v", err)
	}
}
