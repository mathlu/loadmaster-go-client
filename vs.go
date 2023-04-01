package lmclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Vss struct {
	VS []Vs
}
type Vs struct {
	Status        string   `json:"Status"`
	Index         int      `json:"Index"`
	Address       string   `json:"VSAddress"`
	Port          string   `json:"VSPort"`
	Layer         int      `json:"Layer"`
	NickName      string   `json:"NickName"`
	Enable        bool     `json:"Enable"`
	SSLReverse    bool     `json:"SSLReverse"`
	SSLReencrypt  bool     `json:"SSLReencrypt"`
	InterceptMode int      `json:"InterceptMode"`
	Intercept     bool     `json:"Intercept"`
	InterceptOpts []string `json:"InterceptOpts"`
	ForceL4       bool     `json:"ForceL4"`
	ForceL7       bool     `json:"ForceL7"`
	Type          string   `json:"VStype"`
	Protocol      string   `json:"Protocol"`
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
		var interceptopts string
		if len(u.InterceptOpts) > 0 {
			interceptopts = strings.Join(u.InterceptOpts, ";")
		}
		return json.Marshal(&struct {
			ApiKey        string `json:"apikey"`
			CMD           string `json:"cmd"`
			Address       string `json:"vs"`
			Port          string `json:"port"`
			NickName      string `json:"NickName"`
			SSLReverse    bool   `json:"SSLReverse"`
			SSLReencrypt  bool   `json:"SSLReencrypt"`
			InterceptMode int    `json:"InterceptMode"`
			Intercept     bool   `json:"Intercept"`
			InterceptOpts string `json:"InterceptOpts"`
			Enable        bool   `json:"Enable"`
			ForceL4       bool   `json:"ForceL4"`
			ForceL7       bool   `json:"ForceL7"`
			Type          string `json:"VStype"`
			Protocol      string `json:"prot"`
		}{
			ApiKey:        u.ApiKey,
			CMD:           u.CMD,
			Address:       u.Address,
			Port:          u.Port,
			NickName:      u.NickName,
			SSLReverse:    u.SSLReverse,
			SSLReencrypt:  u.SSLReencrypt,
			InterceptMode: u.InterceptMode,
			Intercept:     u.Intercept,
			InterceptOpts: interceptopts,
			Enable:        u.Enable,
			ForceL4:       u.ForceL4,
			ForceL7:       u.ForceL7,
			Type:          u.Type,
			Protocol:      u.Protocol,
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
		var interceptopts string
		if len(u.InterceptOpts) > 0 {
			interceptopts = strings.Join(u.InterceptOpts, ";")
		}
		return json.Marshal(&struct {
			ApiKey        string `json:"apikey"`
			CMD           string `json:"cmd"`
			Index         int    `json:"vs"`
			Address       string `json:"vsaddress"`
			Port          string `json:"vsport"`
			NickName      string `json:"NickName"`
			SSLReverse    bool   `json:"SSLReverse"`
			SSLReencrypt  bool   `json:"SSLReencrypt"`
			InterceptMode int    `json:"InterceptMode"`
			Intercept     bool   `json:"Intercept"`
			InterceptOpts string `json:"InterceptOpts"`
			Enable        bool   `json:"Enable"`
			ForceL4       bool   `json:"ForceL4"`
			ForceL7       bool   `json:"ForceL7"`
			Type          string `json:"VStype"`
			Protocol      string `json:"prot"`
		}{
			ApiKey:        u.ApiKey,
			CMD:           u.CMD,
			Index:         u.Index,
			Address:       u.Address,
			Port:          u.Port,
			NickName:      u.NickName,
			SSLReverse:    u.SSLReverse,
			SSLReencrypt:  u.SSLReencrypt,
			InterceptMode: u.InterceptMode,
			Intercept:     u.Intercept,
			InterceptOpts: interceptopts,
			Enable:        u.Enable,
			ForceL4:       u.ForceL4,
			ForceL7:       u.ForceL7,
			Type:          u.Type,
			Protocol:      u.Protocol,
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
	vsa := &VsApiPayLoad{
		Vs{
			Index: index,
		},
		ApiPayLoad{
			CMD:    "showvs",
			ApiKey: c.ApiKey,
		},
	}
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
	vsa := &VsApiPayLoad{
		Vs{
			Address:       v.Address,
			Port:          v.Port,
			NickName:      v.NickName,
			SSLReverse:    v.SSLReverse,
			SSLReencrypt:  v.SSLReencrypt,
			InterceptMode: v.InterceptMode,
			Intercept:     v.Intercept,
			InterceptOpts: v.InterceptOpts,
			Enable:        v.Enable,
			ForceL4:       v.ForceL4,
			ForceL7:       v.ForceL7,
			Type:          v.Type,
			Protocol:      v.Protocol,
		},
		ApiPayLoad{
			ApiKey: c.ApiKey,
			CMD:    "addvs",
		},
	}

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
	vsa := &VsApiPayLoad{
		Vs{
			Index: index,
		},
		ApiPayLoad{
			CMD:    "delvs",
			ApiKey: c.ApiKey,
		},
	}
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
	vsa := &VsApiPayLoad{
		Vs{
			Index:         v.Index,
			Address:       v.Address,
			Port:          v.Port,
			NickName:      v.NickName,
			SSLReverse:    v.SSLReverse,
			SSLReencrypt:  v.SSLReencrypt,
			InterceptMode: v.InterceptMode,
			Intercept:     v.Intercept,
			InterceptOpts: v.InterceptOpts,
			Enable:        v.Enable,
			ForceL4:       v.ForceL4,
			ForceL7:       v.ForceL7,
			Type:          v.Type,
			Protocol:      v.Protocol,
		},
		ApiPayLoad{
			ApiKey: c.ApiKey,
			CMD:    "modvs",
		},
	}

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
