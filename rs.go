package lmclient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

type Rss struct {
	XMLName xml.Name `xml:"Response"`
	RS      []Rs     `xml:"Success>Data>Rs"`
}

type Rs struct {
	XMLName   xml.Name `xml:"Rs"`
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
	Nrules    int
}

func (c *Client) CreateRs(r *Rs) (*Rs, error) {
	cmd := "addrs"
	rsa := struct {
		ApiKey  string `json:"apikey" qs:"apikey"`
		ApiUser string `json:"apiuser,omitempty" qs:"-"`
		ApiPass string `json:"apipass,omitempty" qs:"-"`
		CMD     string `json:"cmd" qs:"-"`
		VSIndex int    `json:"vs" qs:"vs"`
		Addr    string `json:"rs,omitempty" qs:"rs"`
		Port    int    `json:"rsport,omitempty" qs:"rsport"`
	}{
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		ApiPass: c.ApiPass,
		CMD:     cmd,
		VSIndex: r.VSIndex,
		Addr:    r.Addr,
		Port:    r.Port,
	}

	req, err := c.newRequest(cmd, rsa)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return &Rs{}, err
	}

	var rss Rss
	if c.Version == 1 {
		reader := bytes.NewReader(resp)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
		err = decoder.Decode(&rss)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(resp, &rss)
		if err != nil {
			return nil, err
		}
	}
	return &rss.RS[0], nil
}

func (c *Client) GetRs(index int, vsindex int) (*Rs, error) {
	cmd := "showrs"
	rsa := struct {
		VSIndex int    `json:"vs" qs:"vs"`
		CMD     string `json:"cmd" qs:"-"`
		ApiKey  string `json:"apikey" qs:"apikey"`
		ApiUser string `json:"apiuser,omitempty" qs:"-"`
		ApiPass string `json:"apipass,omitempty" qs:"-"`
		Rsi     string `json:"rs" qs:"rs"`
	}{
		VSIndex: index,
		CMD:     cmd,
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		ApiPass: c.ApiPass,
		Rsi:     "!" + strconv.Itoa(index),
	}
	req, err := c.newRequest(cmd, rsa)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rss Rss
	if c.Version == 1 {
		reader := bytes.NewReader(resp)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
		err = decoder.Decode(&rss)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(resp, &rss)
		if err != nil {
			return nil, err
		}
	}

	return &rss.RS[0], nil
}

func (c *Client) DeleteRs(index int, vsindex int) (*ApiResponse, error) {
	cmd := "delrs"
	rsa := struct {
		VSIndex int    `json:"vs" qs:"vs"`
		CMD     string `json:"cmd" qs:"-"`
		ApiKey  string `json:"apikey" qs:"apikey"`
		ApiUser string `json:"apiuser,omitempty" qs:"-"`
		ApiPass string `json:"apipass,omitempty" qs:"-"`
		Rsi     string `json:"rs" qs:"rs"`
	}{
		VSIndex: index,
		CMD:     cmd,
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		ApiPass: c.ApiPass,
		Rsi:     "!" + strconv.Itoa(index),
	}
	req, err := c.newRequest(cmd, rsa)
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

func (c *Client) ModifyRs(r *Rs) (*ApiResponse, error) {
	cmd := "modrs"
	rsa := struct {
		ApiKey  string `json:"apikey" qs:"apikey"`
		ApiUser string `json:"apiuser,omitempty" qs:"-"`
		ApiPass string `json:"apipass,omitempty" qs:"-"`
		CMD     string `json:"cmd" qs:"-"`
		VSIndex int    `json:"vs" qs:"vs"`
		Rsi     string `json:"rs" qs:"rs"`
		NewPort string `json:"newport" qs:"newport,omitempty"`
	}{
		ApiKey:  c.ApiKey,
		ApiUser: c.ApiUser,
		ApiPass: c.ApiPass,
		CMD:     cmd,
		VSIndex: r.VSIndex,
		Rsi:     "!" + strconv.Itoa(r.RsIndex),
		NewPort: r.NewPort,
	}

	req, err := c.newRequest(cmd, rsa)
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
