package lmclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllVs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/listvs.json")
	ok(t, err)
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/accessv2")
		// Send response to be tested
		_, err := rw.Write([]byte(content))
		if err != nil {
			fmt.Printf("Write failed: %v", err)
		}
	}))

	defer server.Close()
	client := Client{server.Client(), "bar", server.URL}

	vs, err := client.GetAllVs()
	ok(t, err)

	equals(t, vs[0].Index, 1)
	equals(t, vs[0].NickName, "foo")

}

func TestGetVs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/showvs.json")
	ok(t, err)
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/accessv2")
		// Send response to be tested
		_, err := rw.Write([]byte(content))
		if err != nil {
			fmt.Printf("Write failed: %v", err)
		}
	}))

	defer server.Close()
	client := Client{server.Client(), "bar", server.URL}

	vs, err := client.GetVs(1)
	ok(t, err)

	equals(t, vs.Index, 1)
	equals(t, vs.NickName, "foo")

}

func TestMarshalJSON(t *testing.T) {
	var vs VsApiPayLoad
	vs.Address = "192.168.1.10"
	vs.Port = "888"
	vs.Protocol = "tcp"

	ret, err := vs.MarshalJSON()
	ok(t, err)
	equals(t, string(ret), "{\"vs\":\"192.168.1.10\",\"port\":\"888\",\"prot\":\"tcp\"}")
}
