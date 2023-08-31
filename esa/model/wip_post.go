package model

import (
	"time"
)

type WipPostBody struct {
	Wip       bool      `json:"wip"`
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WipPost struct {
	Post *WipPostBody `json:"post"`
}

type WipPostResponse struct {
	URL string `json:"url"`
}
