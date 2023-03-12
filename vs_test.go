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

	equals(t, vs[0].Status, "Down")
	equals(t, vs[0].Index, 1)
	equals(t, vs[0].NickName, "foo")

}
