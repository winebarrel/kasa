package model

type MovePostBody struct {
	Name     string `json:"name,omitempty"`
	Category string `json:"category"`
}

type MovePost struct {
	Post MovePostBody `json:"post"`
}
