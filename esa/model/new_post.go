package model

type NewPostBody struct {
	Name     string   `json:"name,omitempty"`
	BodyMd   string   `json:"body_md,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Category string   `json:"category,omitempty"`
	Wip      *bool    `json:"wip,omitempty"`
	Message  string   `json:"message,omitempty"`
}

type NewPost struct {
	Post NewPostBody `json:"post"`
}

type NewPostResponse struct {
	URL string `json:"url"`
}
