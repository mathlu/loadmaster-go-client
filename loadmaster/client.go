package lmclient

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
        "encoding/json"
        "bytes"
)

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	RestUrl    string
}

type PayLoad struct {
  apikey string
  cmd string
}

func NewClient(apiKey string, restUrl string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
		RestUrl:    restUrl,
	}
}

func (c *Client) newRequest(cmd string) (*http.Request, error) {
        payload := PayLoad{
                apikey: c.ApiKey,
                cmd: cmd,
        }
        b,err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
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
