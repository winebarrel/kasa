package model

import (
	"time"
)

type TagPostBody struct {
	Tags      []string  `json:"tags"`
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TagPost struct {
	Post *TagPostBody `json:"post"`
}

type TagPostResponse struct {
	URL string `json:"url"`
}
