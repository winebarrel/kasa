package model

type Tags struct {
	Tags []struct {
		Name       string `json:"name"`
		PostsCount int    `json:"posts_count"`
	} `json:"tags"`
	PrevPage   *int `json:"prev_page"`
	NextPage   *int `json:"next_page"`
	TotalCount int  `json:"total_count"`
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	MaxPerPage int  `json:"max_per_page"`
}
