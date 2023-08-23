package lmclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestGetRs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/showrs.json")
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

	rs, err := client.GetRs(1, 1)
	ok(t, err)

	equals(t, rs.RsIndex, 1)
	equals(t, rs.Addr, "10.10.10.10")

}

func TestDelRs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/delrs.json")
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

	ar, err := client.DeleteRs(1, 1)
	ok(t, err)

	equals(t, ar.Status, "ok")
	equals(t, ar.Message, "Command completed ok")
	equals(t, ar.Code, 200)

}

func TestCreateRs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/addrs.json")
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

	r := &Rs{
		VSIndex: 1,
		Addr:    "10.10.10.10",
		Port:    8080,
	}

	rs, err := client.CreateRs(r)
	ok(t, err)

	equals(t, rs.Addr, "10.10.10.10")
	equals(t, rs.Port, 8080)

}

func TestCreateRsIntegration(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skipping integration tests, set environment variable INTEGRATION")
	}

	if v := os.Getenv("LOADMASTER_SERVER"); v == "" {
		t.Fatal("LOADMASTER_SERVER must be set for integration tests")
	}

	if v := os.Getenv("LOADMASTER_API_KEY"); v == "" {
		t.Fatal("LOADMASTER_API_KEY must be set for integration tests")
	}
	client := NewClient(os.Getenv("LOADMASTER_API_KEY"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")))
	if vi, err := client.GetVsByName("IntegrationTest"); err == nil {
		_, err := client.DeleteVs(vi.Index)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}
	vs := &Vs{
		Address:  "192.168.1.124",
		Port:     "8093",
		NickName: "IntegrationTest",
		Enable:   true,
		ForceL7:  true,
		Type:     "gen",
		Protocol: "tcp",
	}
	vsc, err := client.CreateVs(vs)
	index := vsc.Index
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	rs := Rs{
		Addr:    "192.168.1.50",
		Port:    80,
		VSIndex: index,
	}
	rsc, err := client.CreateRs(&rs)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	t.Cleanup(func() {
		_, err = client.DeleteRs(rsc.RsIndex, index)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
		_, err = client.DeleteVs(index)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
}

func TestModifyRs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/modrs.json")
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
	r := &Rs{
		VSIndex: 1,
		Rsi:     "!1",
		Addr:    "192.168.1.215",
		NewPort: "6443",
	}

	rs, err := client.ModifyRs(r)
	ok(t, err)

	equals(t, rs.Status, "ok")
	equals(t, rs.Message, "Command completed ok")
	equals(t, rs.Code, 200)

}

func TestModifyRsIntegration(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skipping integration tests, set environment variable INTEGRATION")
	}

	if v := os.Getenv("LOADMASTER_SERVER"); v == "" {
		t.Fatal("LOADMASTER_SERVER must be set for integration tests")
	}

	if v := os.Getenv("LOADMASTER_API_KEY"); v == "" {
		t.Fatal("LOADMASTER_API_KEY must be set for integration tests")
	}
	client := NewClient(os.Getenv("LOADMASTER_API_KEY"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")))
	if vi, err := client.GetVsByName("IntegrationTest"); err == nil {
		_, err := client.DeleteVs(vi.Index)
		ok(t, err)
	}
	vs := &Vs{
		Address:  "192.168.1.123",
		Port:     "8090",
		NickName: "IntegrationTest",
		Enable:   true,
		ForceL7:  true,
		Type:     "gen",
		Protocol: "tcp",
	}
	cvs, err := client.CreateVs(vs)
	ok(t, err)
	index := cvs.Index
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	rs := Rs{
		Addr:    "192.168.1.50",
		Port:    80,
		VSIndex: index,
	}
	rsc, err := client.CreateRs(&rs)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	rsm := Rs{
		Addr:    "192.168.1.50",
		NewPort: "8080",
		VSIndex: index,
		Rsi:     "!" + strconv.Itoa(rsc.RsIndex),
	}

	_, err = client.ModifyRs(&rsm)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	t.Cleanup(func() {
		_, err := client.DeleteVs(index)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
}

func TestGetRsIntegration(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skipping integration tests, set environment variable INTEGRATION")
	}

	if v := os.Getenv("LOADMASTER_SERVER"); v == "" {
		t.Fatal("LOADMASTER_SERVER must be set for integration tests")
	}

	if v := os.Getenv("LOADMASTER_API_KEY"); v == "" {
		t.Fatal("LOADMASTER_API_KEY must be set for integration tests")
	}
	client := NewClient(os.Getenv("LOADMASTER_API_KEY"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")))
	if vi, err := client.GetVsByName("IntegrationTest"); err == nil {
		_, err := client.DeleteVs(vi.Index)
		ok(t, err)
	}
	vs := &Vs{
		Address:  "192.168.1.112",
		Port:     "80",
		NickName: "IntegrationTest",
		Enable:   true,
		ForceL7:  true,
		Type:     "gen",
		Protocol: "tcp",
	}
	cvs, err := client.CreateVs(vs)
	ok(t, err)

	rs := Rs{
		Addr:    "192.168.1.50",
		Port:    80,
		VSIndex: cvs.Index,
	}

	rsc, err := client.CreateRs(&rs)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	grs, err := client.GetRs(rsc.RsIndex, cvs.Index)
	ok(t, err)

	t.Cleanup(func() {
		_, err := client.DeleteVs(grs.VSIndex)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
}
