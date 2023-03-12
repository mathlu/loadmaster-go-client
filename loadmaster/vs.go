package lmclient

import "encoding/json"

type Vss struct {
	code   int
	VS     []Vs
	status string
}
type Vs struct {
	Status   string
	Index    int
	NickName string
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
