package lmclient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Vss struct {
	XMLName xml.Name   `xml:"Response"`
	VS      []VsListed `xml:"Success>Data>VS"`
}

type VsListed struct {
	XMLName  xml.Name `xml:"VS"`
	Index    int      `xml:"Index"`
	NickName string   `xml:"NickName"`
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
	Enable   bool     `xml:"Success>Data>Enable"`
	ForceL4  bool     `xml:"Success>Data>ForceL4"`
	ForceL7  bool     `xml:"Success>Data>ForceL7"`
	Layer    int      `xml:"Success>Data>Layer"`
}

func (c *Client) GetAllVs() ([]VsListed, error) {
	cmd := "listvs"
	payload := struct {
		CMD     string `json:"cmd" qs:"-"`
		ApiKey  string `json:"apikey,omitempty" qs:"apikey,omitempty"`
		ApiUser string `json:"apiuser,omitempty" qs:"-"`
		ApiPass string `json:"apipass,omitempty" qs:"-"`
	}{
		CMD:     cmd,
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		ApiPass: c.ApiPass,
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
		Index   int    `json:"vs" qs:"vs"`
		CMD     string `json:"cmd" qs:"-"`
		ApiKey  string `json:"apikey,omitempty" qs:"apikey,omitempty"`
		ApiUser string `json:"apiuser,omitempty" qs:"-"`
		ApiPass string `json:"apipass,omitempty" qs:"-"`
	}{
		Index:   index,
		CMD:     cmd,
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		ApiPass: c.ApiPass,
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
		content := string(resp)
		re := regexp.MustCompile(">Y<")
		replace_true := re.ReplaceAllString(content, ">true<")
		re = regexp.MustCompile(">N<")
		replaced := re.ReplaceAllString(replace_true, ">false<")

		b := []byte(replaced)

		reader := bytes.NewReader(b)
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

	var enable string
	if v.Enable {
		enable = "Y"
	} else {
		enable = "N"
	}

	var forcel4 int
	var forcel7 int
	if v.Layer == 4 {
		forcel4 = 1
	} else {

		forcel7 = 1
	}

	vsa := struct {
		ApiKey   string `json:"apikey,omitempty" qs:"apikey,omitempty"`
		ApiUser  string `json:"apiuser,omitempty" qs:"-"`
		ApiPass  string `json:"apipass,omitempty" qs:"-"`
		CMD      string `json:"cmd" qs:"-"`
		Address  string `json:"vs" qs:"vs"`
		Port     string `json:"port" qs:"port"`
		NickName string `json:"NickName,omitempty" qs:"nickname,omitempty"`
		Type     string `json:"VStype,omitempty" qs:"vstype,omitempty"`
		Protocol string `json:"prot,omitempty" qs:"prot,omitempty"`
		Enable   string `json:"Enable" qs:"Enable"`
		ForceL4  int    `json:"ForceL4,omitempty" qs:"forcel4"`
		ForceL7  int    `json:"ForceL7,omitempty" qs:"forcel7"`
	}{
		ApiKey:   c.ApiKey,
		ApiUser:  c.ApiUser,
		ApiPass:  c.ApiPass,
		CMD:      cmd,
		Address:  v.Address,
		Port:     v.Port,
		NickName: v.NickName,
		Type:     v.Type,
		Protocol: v.Protocol,
		Enable:   enable,
		ForceL4:  forcel4,
		ForceL7:  forcel7,
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
		content := string(resp)
		re := regexp.MustCompile(">Y<")
		replace_true := re.ReplaceAllString(content, ">true<")
		re = regexp.MustCompile(">N<")
		replaced := re.ReplaceAllString(replace_true, ">false<")

		b := []byte(replaced)

		reader := bytes.NewReader(b)
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
		Index   int    `json:"vs" qs:"vs"`
		CMD     string `json:"cmd" qs:"-"`
		ApiKey  string `json:"apikey,omitempty" qs:"apikey,omitempty"`
		ApiUser string `json:"apiuser,omitempty" qs:"-"`
		ApiPass string `json:"apipass,omitempty" qs:"-"`
	}{
		Index:   index,
		CMD:     cmd,
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		ApiPass: c.ApiPass,
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

	var enable string
	if v.Enable {
		enable = "Y"
	} else {
		enable = "N"
	}

	var forcel4 int
	var forcel7 int
	if v.Layer == 4 {
		forcel4 = 1
	} else {

		forcel7 = 1
	}

	vsa := struct {
		Index    int    `json:"vs" qs:"vs"`
		CMD      string `json:"cmd" qs:"-"`
		ApiKey   string `json:"apikey,omitempty" qs:"apikey,omitempty"`
		ApiUser  string `json:"apiuser,omitempty" qs:"-"`
		ApiPass  string `json:"apipass,omitempty" qs:"-"`
		Address  string `json:"vsaddress" qs:"vsaddress"`
		Port     string `json:"port" qs:"port"`
		VSPort   int    `json:"vsport" qs:"vsport,omitempty"`
		NickName string `json:"NickName,omitempty" qs:"NickName,omitempty"`
		Type     string `json:"VStype,omitempty" qs:"VSType,omitempty"`
		Protocol string `json:"prot,omitempty" qs:"prot,omitempty"`
		Enable   string `json:"Enable" qs:"Enable"`
		ForceL4  int    `json:"ForceL4,omitempty" qs:"forcel4"`
		ForceL7  int    `json:"ForceL7,omitempty" qs:"forcel7"`
	}{
		Index:    v.Index,
		CMD:      cmd,
		ApiKey:   c.ApiKey,
		ApiUser:  c.ApiUser,
		ApiPass:  c.ApiPass,
		Address:  v.Address,
		Port:     v.Port,
		VSPort:   vsport,
		NickName: v.NickName,
		Type:     v.Type,
		Protocol: v.Protocol,
		Enable:   enable,
		ForceL4:  forcel4,
		ForceL7:  forcel7,
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
		content := string(resp)
		re := regexp.MustCompile(">Y<")
		replace_true := re.ReplaceAllString(content, ">true<")
		re = regexp.MustCompile(">N<")
		replaced := re.ReplaceAllString(replace_true, ">false<")

		b := []byte(replaced)

		reader := bytes.NewReader(b)
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
