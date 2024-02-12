package model

import (
	"encoding/json"
	"fmt"
	pathpkg "path"
	"time"

	"github.com/winebarrel/kasa/utils"
)

type PostAuthor struct {
	Myself     bool   `json:"myself"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type Post struct {
	Number         int         `json:"number"`
	Name           string      `json:"name"`
	FullName       string      `json:"full_name,omitempty"`
	Wip            bool        `json:"wip"`
	BodyMd         string      `json:"body_md,omitempty"`
	BodyHTML       string      `json:"body_html,omitempty"`
	CreatedAt      time.Time   `json:"created_at"`
	Message        string      `json:"message"`
	URL            string      `json:"url"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Tags           []string    `json:"tags"`
	Category       string      `json:"category"`
	RevisionNumber int         `json:"revision_number"`
	CreatedBy      *PostAuthor `json:"created_by,omitempty"`
	UpdatedBy      *PostAuthor `json:"updated_by,omitempty"`
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

	urlDir := pathpkg.Dir(post.URL)

	return fmt.Sprintf("%s  %-3s  %-*s  %s  %s",
		post.UpdatedAt.Format("2006-01-02 03:04:05"), wip, len(urlDir)+9, post.URL, post.FullNameWithoutTags(), utils.TagsToString(post.Tags))
}

func (post *Post) Json() (string, error) {
	post.BodyMd = ""
	post.BodyHTML = ""
	post.FullName = ""
	post.CreatedBy = nil
	post.UpdatedBy = nil
	out, err := json.Marshal(post)

	if err != nil {
		return "", err
	}

	return string(out), nil
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
