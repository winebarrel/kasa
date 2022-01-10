package esa

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_SendOK(t *testing.T) {
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

	params := map[string]string{"q": "phrase"}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	client := newClient("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	req, err := client.newRequest(http.MethodGet, "posts", nil)
	assert.NoError(err)

	query := req.URL.Query()
	query.Add("q", "phrase")
	req.URL.RawQuery = query.Encode()

	body, err := client.send(req)
	assert.NoError(err)
	assert.Equal(resBody, string(body))
}

func TestClientSend_NotFound(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	params := map[string]string{"q": "phrase"}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusNotFound, `{"error":"not_found","message":"Not found"}`))

	client := newClient("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", false)
	req, err := client.newRequest(http.MethodGet, "posts", nil)
	assert.NoError(err)

	query := req.URL.Query()
	query.Add("q", "phrase")
	req.URL.RawQuery = query.Encode()

	_, err = client.send(req)
	assert.Equal(errors.New(`404: {"error":"not_found","message":"Not found"}`), err)
}
