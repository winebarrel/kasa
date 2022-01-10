package model

import (
	"strings"
)

type MovePostBody struct {
	Name     string `json:"name,omitempty"`
	Category string `json:"category"`
}

func (post *MovePostBody) String() string {
	cat := strings.TrimSuffix(post.Category, "/")
	return strings.TrimPrefix(cat+"/"+post.Name, "/")
}

type MovePost struct {
	Post MovePostBody `json:"post"`
}
