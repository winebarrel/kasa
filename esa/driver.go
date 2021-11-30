package esa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/winebarrel/kasa/esa/model"
)

const (
	MaxPerPage = 50
)

type Driver struct {
	esaCli *Client
}

func NewDriver(team string, token string) *Driver {
	return &Driver{
		esaCli: newClient(team, token),
	}
}

func (dri *Driver) Get(rawPath string) (*model.Post, error) {
	path := NewPath(rawPath)
	req, err := dri.esaCli.newRequest(http.MethodGet, "posts", nil)

	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("q", path.Name+` on:"`+path.Category+`"`)
	page, err := dri.ListPostsInPage(req, 1, query)

	if err != nil {
		return nil, err
	}

	for _, post := range page.Posts {
		if post.Name == path.Name {
			return post, nil
		}
	}

	return nil, nil
}

func (dri *Driver) List(path *GlobPath, pageNum int, recursive bool) ([]*model.Post, bool, error) {
	req, err := dri.esaCli.newRequest(http.MethodGet, "posts", nil)

	if err != nil {
		return nil, false, err
	}

	queryString := ""

	if path.filename != "" {
		queryString = path.filename
	}

	if recursive {
		queryString += ` in:"` + path.category + `"`
	} else {
		queryString += ` on:"` + path.category + `"`
	}

	query := req.URL.Query()
	query.Add("q", queryString)
	page, err := dri.ListPostsInPage(req, pageNum, query)

	if err != nil {
		return nil, false, err
	}

	posts := []*model.Post{}

	for _, v := range page.Posts {
		if path.match(v.FullNameWithoutTags()) {
			posts = append(posts, v)
		}
	}

	if len(posts) == 0 {
		return nil, false, fmt.Errorf("post not found on page %d", pageNum)
	}

	return posts, page.NextPage != nil, nil
}

func (dri *Driver) Search(queryString string, pageNum int) ([]*model.Post, bool, error) {
	req, err := dri.esaCli.newRequest(http.MethodGet, "posts", nil)

	if err != nil {
		return nil, false, err
	}

	query := req.URL.Query()
	query.Add("q", queryString)
	page, err := dri.ListPostsInPage(req, pageNum, query)

	if err != nil {
		return nil, false, err
	}

	if len(page.Posts) == 0 {
		return nil, false, fmt.Errorf("post not found on page %d", pageNum)
	}

	return page.Posts, page.NextPage != nil, nil
}

func (dri *Driver) ListOrTagSearch(rawPath string, pageNum int, recursive bool) ([]*model.Post, bool, error) {
	if strings.HasPrefix(rawPath, "#") {
		return dri.Search(rawPath, pageNum)
	} else {
		path := newGlobPath(rawPath)
		return dri.List(path, pageNum, recursive)
	}
}

func (dri *Driver) ListPostsInPage(req *http.Request, pageNum int, query url.Values) (*model.Posts, error) {
	query.Add("page", strconv.Itoa(pageNum))
	query.Add("per_page", strconv.Itoa(MaxPerPage))
	req.URL.RawQuery = query.Encode()
	body, err := dri.esaCli.send(req)

	if err != nil {
		return nil, err
	}

	page := &model.Posts{}
	err = json.Unmarshal(body, &page)

	if err != nil {
		return nil, err
	}

	return page, nil
}

func (dri *Driver) Post(newPostBody *model.NewPostBody, postNum int) (string, error) {
	newPost := model.NewPost{
		Post: *newPostBody,
	}

	postBody, err := json.Marshal(newPost)

	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(postBody)
	var req *http.Request

	if postNum > 0 {
		path := fmt.Sprintf("posts/%d", postNum)
		req, err = dri.esaCli.newRequest(http.MethodPatch, path, reader)
	} else {
		req, err = dri.esaCli.newRequest(http.MethodPost, "posts", reader)
	}

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	body, err := dri.esaCli.send(req)

	if err != nil {
		return "", err
	}

	res := model.NewPostResponse{}
	err = json.Unmarshal(body, &res)

	if err != nil {
		return "", err
	}

	return res.URL, nil
}

func (dri *Driver) Move(movePostBody *model.MovePostBody, postNum int) error {
	movePost := model.MovePost{
		Post: *movePostBody,
	}

	postBody, err := json.Marshal(movePost)

	if err != nil {
		return err
	}

	reader := bytes.NewReader(postBody)
	path := fmt.Sprintf("posts/%d", postNum)
	req, err := dri.esaCli.newRequest(http.MethodPatch, path, reader)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	_, err = dri.esaCli.send(req)

	return err
}

func (dri *Driver) MoveCategory(from string, to string) error {
	postBody, err := json.Marshal(&model.MoveCategory{
		From: from,
		To:   to,
	})

	if err != nil {
		return err
	}

	reader := bytes.NewReader(postBody)
	req, err := dri.esaCli.newRequest(http.MethodPost, "categories/batch_move", reader)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	_, err = dri.esaCli.send(req)

	return err
}

func (dri *Driver) Delete(postNum int) error {
	path := fmt.Sprintf("posts/%d", postNum)
	req, err := dri.esaCli.newRequest(http.MethodDelete, path, nil)

	if err != nil {
		return err
	}

	_, err = dri.esaCli.send(req)

	return err
}
