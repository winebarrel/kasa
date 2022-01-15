// cf. https://docs.esa.io/posts/102
package esa

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

const (
	Endpoint = "api.esa.io"
)

type Client struct {
	team    string
	token   string
	httpCli *http.Client
	debug   bool
	version string
}

func newClient(team string, token string, debug bool, version string) *Client {
	return &Client{
		team:    team,
		token:   token,
		httpCli: &http.Client{},
		debug:   debug,
		version: version,
	}
}

func (cli *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("https://%s/v1/teams/%s/%s", Endpoint, cli.team, path)
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+cli.token)
	req.Header.Add("User-Agent", "kasa v"+cli.version)

	return req, nil
}

func (cli *Client) send(req *http.Request) ([]byte, error) {
	if cli.debug {
		b, _ := httputil.DumpRequest(req, true)
		fmt.Printf("---request begin---\n%s\n---request end---\n", b)
	}

	res, err := cli.httpCli.Do(req)

	if cli.debug {
		b, _ := httputil.DumpResponse(res, true)
		fmt.Printf("---response begin---\n%s\n---response end---\n", b)
	}

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
