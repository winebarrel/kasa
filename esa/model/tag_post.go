package model

type TagPostBody struct {
	Tags []string `json:"tags"`
}

type TagPost struct {
	Post *TagPostBody `json:"post"`
}

type TagPostResponse struct {
	URL string `json:"url"`
}
