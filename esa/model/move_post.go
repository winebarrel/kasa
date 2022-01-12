package model

type MovePostBody struct {
	Name     string `json:"name,omitempty"`
	Category string `json:"category"`
	Message  string `json:"message,omitempty"`
}

type MovePost struct {
	Post *MovePostBody `json:"post"`
}
