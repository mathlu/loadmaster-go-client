package lmclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Rss struct {
	RS []Rs
}

type Rs struct {
	Status    string
	Code      int
	Message   string
	VSIndex   int
	RsIndex   int
	Rsi       string
	Addr      string
	Port      int
	NewPort   string
	DnsName   string
	Forward   string
	Weight    int
	Limit     int
	RateLimit int
	Follow    int
	Enable    bool
	Critical  bool
	Nrules    int
}

type RsApiPayLoad struct {
	*Rs
	ApiPayLoad
}

func (u RsApiPayLoad) MarshalJSON() ([]byte, error) {
	switch u.CMD {
	case "addrs":
		return json.Marshal(&struct {
			ApiKey  string `json:"apikey"`
			CMD     string `json:"cmd"`
			VSIndex int    `json:"vs"`
			RsIndex int    `json:"RsIndex,omitempty"`
			Addr    string `json:"rs,omitempty"`
			Port    int    `json:"rsport,omitempty"`
			NewPort string `json:"newport,omitempty"`
		}{
			ApiKey:  u.ApiKey,
			CMD:     u.CMD,
			VSIndex: u.VSIndex,
			RsIndex: u.RsIndex,
			Addr:    u.Addr,
			Port:    u.Port,
			NewPort: u.NewPort,
		})
	case "modrs":
		return json.Marshal(&struct {
			ApiKey  string `json:"apikey"`
			CMD     string `json:"cmd"`
			VSIndex int    `json:"vs"`
			Rsi     string `json:"rs"`
			NewPort string `json:"newport"`
		}{
			ApiKey:  u.ApiKey,
			CMD:     u.CMD,
			VSIndex: u.VSIndex,
			Rsi:     u.Rsi,
			NewPort: u.NewPort,
		})
	case "showrs", "delrs":
		return json.Marshal(&struct {
			ApiKey  string `json:"apikey"`
			CMD     string `json:"cmd"`
			VSIndex int    `json:"vs"`
			Rsi     string `json:"rs"`
		}{
			ApiKey:  u.ApiKey,
			CMD:     u.CMD,
			VSIndex: u.VSIndex,
			Rsi:     u.Rsi,
		})
	default:
		return nil, errors.New("Unknown CMD")

	}

}

func (c *Client) CreateRs(r *Rs) (*Rs, error) {
	rsa := &RsApiPayLoad{
		r,
		ApiPayLoad{
			ApiKey: c.ApiKey,
			CMD:    "addrs",
		},
	}

	b, err := json.Marshal(rsa)
	if err != nil {
		return &Rs{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return &Rs{}, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return &Rs{}, err
	}

	var rss Rss
	err = json.Unmarshal(resp, &rss)
	if err != nil {
		return nil, err
	}

	return &rss.RS[0], nil
}

func (c *Client) GetRs(index int, vsindex int) (*Rs, error) {
	rsa := &RsApiPayLoad{
		&Rs{
			VSIndex: vsindex,
			Rsi:     "!" + strconv.Itoa(index),
		},
		ApiPayLoad{
			CMD:    "showrs",
			ApiKey: c.ApiKey,
		},
	}
	b, err := json.Marshal(rsa)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rss Rss
	err = json.Unmarshal(resp, &rss)
	if err != nil {
		return nil, err
	}

	return &rss.RS[0], nil
}

func (c *Client) DeleteRs(index int, vsindex int) (*ApiResponse, error) {
	rsa := &RsApiPayLoad{
		&Rs{
			VSIndex: vsindex,
			Rsi:     "!" + strconv.Itoa(index),
		},
		ApiPayLoad{
			CMD:    "delrs",
			ApiKey: c.ApiKey,
		},
	}
	b, err := json.Marshal(rsa)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var ar ApiResponse
	err = json.Unmarshal(resp, &ar)
	if err != nil {
		return nil, err
	}

	if ar.Status != "ok" {
		return nil, errors.New("Code: " + fmt.Sprint(ar.Code) + " Message:" + ar.Message)
	}

	return &ar, nil
}

func (c *Client) ModifyRs(r *Rs) (*Rs, error) {
	rsa := &RsApiPayLoad{
		r,
		ApiPayLoad{
			ApiKey: c.ApiKey,
			CMD:    "modrs",
		},
	}

	b, err := json.Marshal(rsa)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rs Rs
	err = json.Unmarshal(resp, &rs)
	if err != nil {
		return nil, err
	}

	return &rs, nil
}
