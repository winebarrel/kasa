package esa_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa/esa"
	"github.com/winebarrel/kasa/esa/model"
)

func TestDriverGetFromPageNum_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"number": 1,
		"name": "hi!",
		"full_name": "日報/2015/05/09/hi! #api #dev",
		"wip": true,
		"body_md": "# Getting Started",
		"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
		"created_at": "2015-05-09T11:54:50+09:00",
		"message": "Add Getting Started section",
		"url": "https://docs.esa.io/posts/1",
		"updated_at": "2015-05-09T11:54:51+09:00",
		"tags": [
			"api",
			"dev"
		],
		"category": "日報/2015/05/09",
		"revision_number": 1,
		"created_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"updated_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"kind": "flow",
		"comments_count": 1,
		"tasks_count": 1,
		"done_tasks_count": 1,
		"stargazers_count": 1,
		"watchers_count": 1,
		"star": true,
		"watch": true
	}`

	httpmock.RegisterResponder(http.MethodGet, "https://api.esa.io/v1/teams/example/posts/1",
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	post, err := driver.GetFromPageNum(1)
	assert.NoError(err)

	expected := &model.Post{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected, post)
}

func TestDriverGet_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBodyPost := `{
		"number": 1,
		"name": "hi!",
		"full_name": "日報/2015/05/09/hi! #api #dev",
		"wip": true,
		"body_md": "# Getting Started",
		"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
		"created_at": "2015-05-09T11:54:50+09:00",
		"message": "Add Getting Started section",
		"url": "https://docs.esa.io/posts/1",
		"updated_at": "2015-05-09T11:54:51+09:00",
		"tags": [
			"api",
			"dev"
		],
		"category": "日報/2015/05/09",
		"revision_number": 1,
		"created_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"updated_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"kind": "flow",
		"comments_count": 1,
		"tasks_count": 1,
		"done_tasks_count": 1,
		"stargazers_count": 1,
		"watchers_count": 1,
		"star": true,
		"watch": true
	}`

	params := map[string]string{"page": "1", "per_page": "50", "q": `hi! on:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, `{"posts":[`+resBodyPost+`]}`))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	post, err := driver.Get("日報/2015/05/09/hi!")
	assert.NoError(err)

	expected := &model.Post{}
	json.Unmarshal([]byte(resBodyPost), expected)
	assert.Equal(expected, post)
}

func TestDriverList_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "2", "per_page": "50", "q": `hi! in:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.List("日報/2015/05/09/hi!", 2, true)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestDriverList_MatchNameOnly(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!!",
				"category": "日報/2015/05/09"
			},
			{
				"number": 2,
				"name": "hi!",
				"category": "日報/2015/05/09"
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "2", "per_page": "50", "q": `hi! in:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.List("日報/2015/05/09/hi!", 2, true)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts[1:2], posts)
	assert.False(hasMore)
}

func TestDriverList_NotFound(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{"posts": []}`

	params := map[string]string{"page": "2", "per_page": "50", "q": `hi! in:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	_, _, err := driver.List("日報/2015/05/09/hi!", 2, true)
	assert.Equal(errors.New("post not found"), err)
}

func TestDriverList_NotFound_HasMore(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{"posts": [], "next_page": 3}`

	params := map[string]string{"page": "2", "per_page": "50", "q": `hi! in:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	_, _, err := driver.List("日報/2015/05/09/hi!", 2, true)
	assert.Equal(errors.New("post not found on page 2"), err)
}

func TestDriverList_WithCategory(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "2", "per_page": "50", "q": ` in:"日報/2015/05/09/"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.List("日報/2015/05/09/", 2, true)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestDriverList_HasMore(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": 2,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "2", "per_page": "50", "q": `hi! in:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.List("日報/2015/05/09/hi!", 2, true)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.True(hasMore)
}

func TestDriverList_NoRecursive(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "3", "per_page": "50", "q": `hi! on:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.List("日報/2015/05/09/hi!", 3, false)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestDriverSearch_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "1", "per_page": "50", "q": `日報`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.Search("日報", 1)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestDriverSearch_NotFound(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{"posts": []}`

	params := map[string]string{"page": "1", "per_page": "50", "q": `日報`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	_, _, err := driver.Search("日報", 1)
	assert.Equal(errors.New("post not found"), err)
}

func TestDriverListOrTagSearch_List(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "1", "per_page": "50", "q": `hi! on:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.ListOrTagSearch("日報/2015/05/09/hi!", 1, false)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestDriverListOrTagSearch_Search(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "1", "per_page": "50", "q": `#tag`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	posts, hasMore, err := driver.ListOrTagSearch("#tag", 1, false)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestDriverPost_Post(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://api.esa.io/v1/teams/example/posts", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"name":"name","body_md":"body_md","tags":["tagA","tagB"],"category":"foo/bar","wip":false,"message":"[skip notice] message"}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusCreated, `{"url":"https://docs.esa.io/posts/5"}`), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	post := &model.NewPostBody{
		Name:     "name",
		BodyMd:   "body_md",
		Tags:     []string{"tagA", "tagB"},
		Category: "foo/bar",
		Wip:      esa.Bool(false),
		Message:  "message",
	}

	url, err := driver.Post(post, 0, false)
	assert.Equal("https://docs.esa.io/posts/5", url)
	assert.NoError(err)
}

func TestDriverPost_WithNotify(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://api.esa.io/v1/teams/example/posts", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"name":"name","body_md":"body_md","tags":["tagA","tagB"],"category":"foo/bar","wip":false,"message":"message"}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusCreated, `{"url":"https://docs.esa.io/posts/5"}`), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	post := &model.NewPostBody{
		Name:     "name",
		BodyMd:   "body_md",
		Tags:     []string{"tagA", "tagB"},
		Category: "foo/bar",
		Wip:      esa.Bool(false),
		Message:  "message",
	}

	url, err := driver.Post(post, 0, true)
	assert.Equal("https://docs.esa.io/posts/5", url)
	assert.NoError(err)
}

func TestDriverPost_Update(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPatch, "https://api.esa.io/v1/teams/example/posts/1", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"name":"name","body_md":"body_md","tags":["tagA","tagB"],"category":"foo/bar","wip":false,"message":"[skip notice] message"}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusOK, `{"url":"https://docs.esa.io/posts/5"}`), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	post := &model.NewPostBody{
		Name:     "name",
		BodyMd:   "body_md",
		Tags:     []string{"tagA", "tagB"},
		Category: "foo/bar",
		Wip:      esa.Bool(false),
		Message:  "message",
	}

	url, err := driver.Post(post, 1, false)
	assert.Equal("https://docs.esa.io/posts/5", url)
	assert.NoError(err)
}

func TestDriverMove_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPatch, "https://api.esa.io/v1/teams/example/posts/1", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"name":"new_name","category":"new_cat","message":"[skip notice]"}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	movePost := &model.MovePostBody{
		Name:     "new_name",
		Category: "new_cat",
	}

	err := driver.Move(movePost, 1, false)
	assert.NoError(err)
}

func TestDriverMove_WithNotify(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPatch, "https://api.esa.io/v1/teams/example/posts/1", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"name":"new_name","category":"new_cat"}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	movePost := &model.MovePostBody{
		Name:     "new_name",
		Category: "new_cat",
	}

	err := driver.Move(movePost, 1, true)
	assert.NoError(err)
}

func TestDriverTag_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPatch, "https://api.esa.io/v1/teams/example/posts/1", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"tags":["foo","bar","zoo"],"message":"[skip notice]"}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	tagPost := &model.TagPostBody{
		Tags: []string{"foo", "bar", "zoo"},
	}

	err := driver.Tag(tagPost, 1, false)
	assert.NoError(err)
}

func TestDriverTag_Delete(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPatch, "https://api.esa.io/v1/teams/example/posts/1", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"tags":[],"message":"[skip notice]"}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	tagPost := &model.TagPostBody{
		Tags: []string{},
	}

	err := driver.Tag(tagPost, 1, false)
	assert.NoError(err)
}

func TestDriverTag_WithNotify(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPatch, "https://api.esa.io/v1/teams/example/posts/1", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"post":{"tags":["foo","bar","zoo"]}}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)

	tagPost := &model.TagPostBody{
		Tags: []string{"foo", "bar", "zoo"},
	}

	err := driver.Tag(tagPost, 1, true)
	assert.NoError(err)
}

func TestDriverMoveCategory_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://api.esa.io/v1/teams/example/categories/batch_move", func(req *http.Request) (*http.Response, error) {
		resBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(`{"from":"from_cat","to":"to_cat"}`, string(resBody))
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	err := driver.MoveCategory("from_cat", "to_cat")
	assert.NoError(err)
}

func TestDriverDriverDelete_Ok(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://api.esa.io/v1/teams/example/posts/1",
		httpmock.NewStringResponder(http.StatusNoContent, ""))

	driver := esa.NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	err := driver.Delete(1)
	assert.NoError(err)
}
