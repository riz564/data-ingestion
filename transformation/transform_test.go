package data_transformation

import (
	"testing"
	"time"

	"data-ingestion/model"
)

func TestTransformPosts(t *testing.T) {
	posts := []model.Post{
		{Id: 1, Title: "Test", Body: "Body", UserId: 42},
	}
	source := "placeholder_api"
	transformed := TransformPosts(posts, source)

	if len(transformed) != 1 {
		t.Fatalf("expected 1 post, got %d", len(transformed))
	}
	got := transformed[0]
	if got.Source != source {
		t.Errorf("expected source %q, got %q", source, got.Source)
	}
	if got.IngestedAt.IsZero() {
		t.Error("expected ingested_at to be set, got zero value")
	}

	if got.Id != posts[0].Id || got.Title != posts[0].Title || got.Body != posts[0].Body || got.UserId != posts[0].UserId {
		t.Error("original post fields not preserved")
	}

	if time.Since(got.IngestedAt) > 2*time.Second {
		t.Error("ingested_at is not recent")
	}
}
