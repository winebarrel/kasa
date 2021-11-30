package model

type NewPostBody struct {
	Name     string   `json:"name"`
	BodyMd   string   `json:"body_md"`
	Tags     []string `json:"tags"`
	Category string   `json:"category"`
	WIP      bool     `json:"wip"`
	Message  string   `json:"message"`
}

type NewPost struct {
	Post NewPostBody `json:"post"`
}

type NewPostResponse struct {
	URL string `json:"url"`
}
