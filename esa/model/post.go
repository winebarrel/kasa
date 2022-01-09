package model

import (
	"fmt"
	pathpkg "path"
	"strings"
	"time"
)

type Post struct {
	Number         int       `json:"number"`
	Name           string    `json:"name"`
	FullName       string    `json:"full_name"`
	Wip            bool      `json:"wip"`
	BodyMd         string    `json:"body_md,omitempty"`
	BodyHTML       string    `json:"body_html,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	Message        string    `json:"message"`
	URL            string    `json:"url"`
	UpdatedAt      time.Time `json:"updated_at"`
	Tags           []string  `json:"tags"`
	Category       string    `json:"category"`
	RevisionNumber int       `json:"revision_number"`
	CreatedBy      struct {
		Myself     bool   `json:"myself"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Icon       string `json:"icon"`
	} `json:"created_by"`
	UpdatedBy struct {
		Myself     bool   `json:"myself"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Icon       string `json:"icon"`
	} `json:"updated_by"`
}

func (post *Post) FullNameWithoutTags() string {
	category := post.Category
	name := post.Name

	if category != "" {
		return pathpkg.Join(category, name)
	} else {
		return name
	}
}

func (post *Post) ListString() string {
	var wip string

	if post.Wip {
		wip = "WIP"
	} else {
		wip = "-"
	}

	var tags string

	if len(post.Tags) > 0 {
		tags = "[#" + strings.Join(post.Tags, ",#") + "]"
	} else {
		tags = ""
	}

	urlDir := pathpkg.Dir(post.URL)
	return fmt.Sprintf("%s  %-3s  %-*s  %s  %s", post.UpdatedAt.Format("2006-01-02 03:04:05"), wip, len(urlDir)+9, post.URL, post.FullNameWithoutTags(), tags)
}

type Posts struct {
	Posts      []*Post `json:"posts"`
	PrevPage   *int    `json:"prev_page"`
	NextPage   *int    `json:"next_page"`
	TotalCount int     `json:"total_count"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	MaxPerPage int     `json:"max_per_page"`
}
