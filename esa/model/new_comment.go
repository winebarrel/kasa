package model

type NewCommentBody struct {
	BodyMd string `json:"body_md"`
}

type NewComment struct {
	Comment *NewCommentBody `json:"comment"`
}

type NewCommentResponse struct {
	URL string `json:"url"`
}
