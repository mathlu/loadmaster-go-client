package lmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Vss struct {
	VS []Vs
}
type Vs struct {
	Status   string `json:"Status"`
	Index    int
	NickName string
}

type GetVsPayLoad struct {
	ApiKey  string `json:"apikey"`
	CMD     string `json:"cmd"`
	VsIndex int    `json:"vs"`
}

func (c *Client) GetAllVs() ([]Vs, error) {
	req, err := c.newRequest("listvs")
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vss Vss
	err = json.Unmarshal(resp, &vss)
	if err != nil {
		return nil, err
	}

	return vss.VS, err
}

func (c *Client) GetVs(index int) (*Vs, error) {
	payload := GetVsPayLoad{
		ApiKey:  c.ApiKey,
		CMD:     "showvs",
		VsIndex: index,
	}
	b, err := json.Marshal(payload)
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

	var vs Vs
	err = json.Unmarshal(resp, &vs)
	if err != nil {
		return nil, err
	}

	return &vs, nil
}
