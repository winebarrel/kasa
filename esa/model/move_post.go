package model

import (
	"time"
)

type MovePostBody struct {
	Name      string    `json:"name,omitempty"`
	Category  string    `json:"category"`
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MovePost struct {
	Post *MovePostBody `json:"post"`
}
