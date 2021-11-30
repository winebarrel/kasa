// cf. https://docs.esa.io/posts/102
package esa

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	Endpoint = "api.esa.io"
)

type Client struct {
	team    string
	token   string
	httpCli *http.Client
}

func newClient(team string, token string) *Client {
	return &Client{
		team:    team,
		token:   token,
		httpCli: &http.Client{},
	}
}

func (cli *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("https://%s/v1/teams/%s/%s", Endpoint, cli.team, path)
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+cli.token)

	return req, nil
}

func (cli *Client) send(req *http.Request) ([]byte, error) {
	res, err := cli.httpCli.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || 299 < res.StatusCode {
		return nil, fmt.Errorf("%s: %s", res.Status, body)
	}

	return body, nil
}
