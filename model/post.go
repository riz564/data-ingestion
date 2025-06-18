package model

import "time"

type Post struct {
	Body       string    `json:"body"`
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	UserId     int       `json:"userId"`
	IngestedAt time.Time `json:"ingested_at,omitempty"`
	Source     string    `json:"source,omitempty"`
}
