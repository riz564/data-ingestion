package data_transformation

import (
	"time"

	"data-ingestion/model"
)

func TransformPosts(posts []model.Post, source string) []model.Post {
	now := time.Now().UTC()
	for i := range posts {
		posts[i].IngestedAt = now
		posts[i].Source = source
	}
	return posts
}
