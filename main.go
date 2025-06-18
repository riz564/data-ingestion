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
		panic(err)
	}
	transformed := transformer.TransformPosts(posts, "placeholder_api")
	
	on:
	  push:
		branches: [ main ]
	  pull_request:
		branches: [ main ]
	
	jobs:
	  build-test:
		runs-on: ubuntu-latest
		steps:
		  - uses: actions/checkout@v4
	
		  - name: Set up Go
			uses: actions/setup-go@v5
			with:
			  go-version: '1.21'
	
		  - name: Install dependencies
			run: go mod download
	
		  - name: Run tests
			run: go test ./...
	
	  deploy:
		needs: build-test
		runs-on: ubuntu-latest
		if: github.ref == 'refs/heads/main'
		steps:
		  - uses: actions/checkout@v4
	
		  - name: Set up Google Cloud SDK
			uses: google-github-actions/setup-gcloud@v2
			with:
			  project_id: ${{ secrets.GCP_PROJECT_ID }}
			  service_account_key: ${{ secrets.GCP_SA_KEY }}
	
		  - name: Build and push Docker image
			run: |
			  gcloud builds submit --tag gcr.io/${{ secrets.GCP_PROJECT_ID }}/data-ingestion
	
		  - name: Deploy to Cloud Run
			run: |
			  gcloud run deploy data-ingestion \
				--image gcr.io/${{ secrets.GCP_PROJECT_ID }}/data-ingestion \
				--region us-central1 \
				--platform managed \
				--allow-unauthenticated
	ctx := context.Background()
	err = storage.StorePosts(ctx, transformed)
	if err != nil {
		log.Fatalf("Failed to store posts: %v", err)
	}

}
