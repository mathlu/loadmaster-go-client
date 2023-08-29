package lmclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pasztorpisti/qs"
)

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	ApiUser    string
	ApiPass    string
	RestUrl    string
	Version    int
}

type ApiResponse struct {
	XMLName xml.Name `xml:"Response"`
	Code    int      `json:"code" xml:"stat,attr"`
	Message string   `json:"message" xml:"Success"`
	Status  string   `json:"status" xml:"code,attr"`
}

func NewClient(apiKey string, apiUser string, apiPass string, restUrl string, version int) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
		ApiUser:    apiUser,
		ApiPass:    apiPass,
		RestUrl:    restUrl,
		Version:    version,
	}
}

func (c *Client) newRequest(cmd string, payload interface{}) (*http.Request, error) {
	if (c.ApiUser == "" || c.ApiPass == "") && c.ApiKey == "" {
		err := fmt.Errorf("Missing authentication")
		return nil, err
	}
	if c.Version == 1 {
		v, _ := qs.Marshal(payload)
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/access/%s?%s", c.RestUrl, cmd, v), nil)
		if err != nil {
			return nil, err
		}
		if c.ApiUser != "" && c.ApiPass != "" {
			req.SetBasicAuth(c.ApiUser, c.ApiPass)
		}
		return req, nil
	} else {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
		if err != nil {
			return nil, err
		}
		return req, nil
	}

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
