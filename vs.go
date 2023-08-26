package lmclient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

type Vss struct {
	XMLName xml.Name   `xml:"Response"`
	VS      []VsListed `xml:"Success>Data>VS"`
}

type VsListed struct {
	XMLName  xml.Name `xml:"VS"`
	Index    int      `xml:"Index"`
	Address  string   `json:"VSAddress" xml:"VSAddress"`
	Port     string   `xml:"Port"`
	VSPort   string   `xml:"VSPort"`
	NickName string   `xml:"NickName"`
	Type     string   `json:"VSType" xml:"VStype"`
	Protocol string   `xml:"Protocol"`
}

type Vs struct {
	XMLName  xml.Name `xml:"Response"`
	Index    int      `xml:"Success>Data>Index"`
	Address  string   `json:"VSAddress" xml:"Success>Data>VSAddress"`
	Port     string   `xml:"Success>Data>Port"`
	VSPort   string   `xml:"Success>Data>VSPort"`
	NickName string   `xml:"Success>Data>NickName"`
	Type     string   `json:"VSType" xml:"Success>Data>VStype"`
	Protocol string   `xml:"Success>Data>Protocol"`
}

func (c *Client) GetAllVs() ([]VsListed, error) {
	cmd := "listvs"
	payload := struct {
		CMD    string `json:"cmd" qs:"-"`
		ApiKey string `json:"apikey" qs:"apikey"`
	}{
		CMD:    cmd,
		ApiKey: c.ApiKey,
	}

	req, err := c.newRequest("listvs", payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var vss Vss
	if c.Version == 1 {
		reader := bytes.NewReader(resp)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
		err = decoder.Decode(&vss)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(resp, &vss)
		if err != nil {
			return nil, err
		}
	}
	return vss.VS, nil

}

func (c *Client) GetVsByName(nickname string) (*VsListed, error) {
	vss, err := c.GetAllVs()
	if err != nil {
		return nil, err
	}

	for _, vs := range vss {
		if vs.NickName == nickname {
			return &vs, nil
		}
	}

	return nil, fmt.Errorf("Virtual Service with name %s not found", nickname)
}

func (c *Client) GetVs(index int) (*Vs, error) {
	cmd := "showvs"
	vsa := struct {
		Index  int    `json:"vs" qs:"vs"`
		CMD    string `json:"cmd" qs:"-"`
		ApiKey string `json:"apikey" qs:"apikey"`
	}{
		Index:  index,
		CMD:    cmd,
		ApiKey: c.ApiKey,
	}

	req, err := c.newRequest(cmd, vsa)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vs Vs
	if c.Version == 1 {
		reader := bytes.NewReader(resp)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
		err = decoder.Decode(&vs)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(resp, &vs)
		if err != nil {
			return nil, err
		}
	}
	return &vs, nil
}

func (c *Client) CreateVs(v *Vs) (*Vs, error) {
	cmd := "addvs"
	vsa := struct {
		ApiKey   string `json:"apikey" qs:"apikey"`
		CMD      string `json:"cmd" qs:"-"`
		Address  string `json:"vs" qs:"vs"`
		Port     string `json:"port" qs:"port"`
		NickName string `json:"NickName,omitempty" qs:"nickname,omitempty"`
		Type     string `json:"VStype,omitempty" qs:"vstype,omitempty"`
		Protocol string `json:"prot,omitempty" qs:"prot,omitempty"`
	}{
		ApiKey:   c.ApiKey,
		CMD:      cmd,
		Address:  v.Address,
		Port:     v.Port,
		NickName: v.NickName,
		Type:     v.Type,
		Protocol: v.Protocol,
	}

	req, err := c.newRequest(cmd, vsa)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vs Vs
	if c.Version == 1 {
		reader := bytes.NewReader(resp)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
		err = decoder.Decode(&vs)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(resp, &vs)
		if err != nil {
			return nil, err
		}
	}
	return &vs, nil

}

func (c *Client) DeleteVs(index int) (*ApiResponse, error) {
	cmd := "delvs"
	vsa := struct {
		Index  int    `json:"vs" qs:"vs"`
		CMD    string `json:"cmd" qs:"-"`
		ApiKey string `json:"apikey" qs:"apikey"`
	}{
		Index:  index,
		CMD:    cmd,
		ApiKey: c.ApiKey,
	}

	req, err := c.newRequest(cmd, vsa)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var ar ApiResponse
	if c.Version == 1 {
		reader := bytes.NewReader(resp)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
		err = decoder.Decode(&ar)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(resp, &ar)
		if err != nil {
			return nil, err
		}
	}

	if ar.Status != "ok" {
		return nil, errors.New("Code: " + fmt.Sprint(ar.Code) + " Message:" + ar.Message)
	}

	return &ar, nil
}

func (c *Client) ModifyVs(v *Vs) (*Vs, error) {
	cmd := "modvs"
	vsport, _ := strconv.Atoi(v.VSPort)
	vsa := struct {
		Index    int    `json:"vs" qs:"vs"`
		CMD      string `json:"cmd" qs:"-"`
		ApiKey   string `json:"apikey" qs:"apikey"`
		Address  string `json:"vsaddress" qs:"vsaddress"`
		Port     string `json:"port" qs:"port"`
		VSPort   int    `json:"vsport" qs:"vsport,omitempty"`
		NickName string `json:"NickName,omitempty" qs:"NickName,omitempty"`
		Type     string `json:"VStype,omitempty" qs:"VSType,omitempty"`
		Protocol string `json:"prot,omitempty" qs:"prot,omitempty"`
	}{
		Index:    v.Index,
		CMD:      cmd,
		ApiKey:   c.ApiKey,
		Address:  v.Address,
		Port:     v.Port,
		VSPort:   vsport,
		NickName: v.NickName,
		Type:     v.Type,
		Protocol: v.Protocol,
	}

	req, err := c.newRequest(cmd, vsa)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vs Vs
	if c.Version == 1 {
		reader := bytes.NewReader(resp)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
		err = decoder.Decode(&vs)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(resp, &vs)
		if err != nil {
			return nil, err
		}
	}

	return &vs, nil
}
