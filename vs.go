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
	Port     string `json:"VSPort"`
	Protocol string `json:"Protocol"`
	Address  string `json:"VSAddress"`
	NickName string `json:"NickName"`
	Layer    int    `json:"Layer"`
	Enable   bool   `json:"Enable"`
	Type     string `json:"VStype"`
	ForceL4  bool   `json:"ForceL4"`
	ForceL7  bool   `json:"ForceL7"`
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
			NickName string `json:"NickName"`
			Enable   bool   `json:"Enable"`
			Type     string `json:"VStype"`
			ForceL4  bool   `json:"ForceL4"`
			ForceL7  bool   `json:"ForceL7"`
		}{
			Address:  u.Address,
			Port:     u.Port,
			Protocol: u.Protocol,
			ApiKey:   u.ApiKey,
			CMD:      u.CMD,
			NickName: u.NickName,
			Enable:   u.Enable,
			Type:     u.Type,
			ForceL4:  u.ForceL4,
			ForceL7:  u.ForceL7,
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
			NickName string `json:"NickName"`
			Enable   bool   `json:"Enable"`
			Type     string `json:"VStype"`
			ForceL4  bool   `json:"ForceL4"`
			ForceL7  bool   `json:"ForceL7"`
		}{
			Index:    u.Index,
			Address:  u.Address,
			Port:     u.Port,
			Protocol: u.Protocol,
			ApiKey:   u.ApiKey,
			CMD:      u.CMD,
			NickName: u.NickName,
			Enable:   u.Enable,
			Type:     u.Type,
			ForceL4:  u.ForceL4,
			ForceL7:  u.ForceL7,
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

func (c *Client) CreateVs(v *Vs) (*Vs, error) {
	var vsa VsApiPayLoad
	vsa.Address = v.Address
	vsa.Port = v.Port
	vsa.Protocol = v.Protocol
	vsa.NickName = v.NickName
	vsa.Enable = v.Enable
	vsa.Type = v.Type
	vsa.ForceL4 = v.ForceL4
	vsa.ForceL7 = v.ForceL7
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
	vsa.CMD = "delvs"
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

func (c *Client) ModifyVs(v *Vs) (*Vs, error) {
	var vsa VsApiPayLoad
	vsa.Index = v.Index
	vsa.Address = v.Address
	vsa.Port = v.Port
	vsa.Protocol = v.Protocol
	vsa.ApiKey = c.ApiKey
	vsa.NickName = v.NickName
	vsa.Enable = v.Enable
	vsa.Type = v.Type
	vsa.ForceL4 = v.ForceL4
	vsa.ForceL7 = v.ForceL7
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
