package lmclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Vss struct {
	VS []Vs
}
type Vs struct {
	Status   string `json:"Status"`
	Index    int    `json:"Index"`
	NickName string `json:"NickName"`
	Port     string `json:"VSPort"`
	Protocol string `json:"Protocol"`
	Address  string `json:"VSAddress"`
}

type ApiPayLoad struct {
	ApiKey string `json:"apikey"`
	CMD    string `json:"cmd"`
}

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type VsApiPayLoad struct {
	Vs
	ApiPayLoad
}

func (u VsApiPayLoad) MarshalJSON() ([]byte, error) {
	switch u.CMD {
	case "addvs":
		return json.Marshal(&struct {
			Address  string `json:"vs"`
			Port     string `json:"port"`
			Protocol string `json:"prot"`
			ApiKey   string `json:"apikey"`
			CMD      string `json:"cmd"`
		}{
			Address:  u.Address,
			Port:     u.Port,
			Protocol: u.Protocol,
			ApiKey:   u.ApiKey,
			CMD:      u.CMD,
		})
	case "delvs", "showvs":
		return json.Marshal(&struct {
			Index  int    `json:"vs"`
			ApiKey string `json:"apikey"`
			CMD    string `json:"cmd"`
		}{
			Index:  u.Index,
			ApiKey: u.ApiKey,
			CMD:    u.CMD,
		})
	case "modvs":
		return json.Marshal(&struct {
			Index    int    `json:"vs"`
			Address  string `json:"vsaddress"`
			Port     string `json:"vsport"`
			Protocol string `json:"prot"`
			ApiKey   string `json:"apikey"`
			CMD      string `json:"cmd"`
		}{
			Index:    u.Index,
			Address:  u.Address,
			Port:     u.Port,
			Protocol: u.Protocol,
			ApiKey:   u.ApiKey,
			CMD:      u.CMD,
		})
	default:
		return nil, errors.New("Unknown CMD")
	}
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
	var vsa VsApiPayLoad
	vsa.CMD = "showvs"
	vsa.ApiKey = c.ApiKey
	vsa.Index = index
	b, err := json.Marshal(vsa)
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

func (c *Client) CreateVs(ip string, proto string, port string) (*Vs, error) {
	var vsa VsApiPayLoad
	vsa.Address = ip
	vsa.Port = port
	vsa.Protocol = proto
	vsa.ApiKey = c.ApiKey
	vsa.CMD = "addvs"

	b, err := json.Marshal(vsa)
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

func (c *Client) DeleteVs(index int) (*ApiResponse, error) {
	var vsa VsApiPayLoad
	vsa.CMD = "showvs"
	vsa.ApiKey = c.ApiKey
	vsa.Index = index
	b, err := json.Marshal(vsa)
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

func (c *Client) ModifyVs(index int, ip string, proto string, port string) (*Vs, error) {
	var vsa VsApiPayLoad
	vsa.Index = index
	vsa.Address = ip
	vsa.Port = port
	vsa.Protocol = proto
	vsa.ApiKey = c.ApiKey
	vsa.CMD = "modvs"

	b, err := json.Marshal(vsa)
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
