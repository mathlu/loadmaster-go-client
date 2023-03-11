package lmclient

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	RestUrl    string
}

func NewClient(apiKey string, restUrl string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
		RestUrl:    restUrl,
	}
}

func (c *Client) newRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s&apikey=%s", c.RestUrl, path, c.ApiKey), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
