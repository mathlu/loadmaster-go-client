package lmclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestDelVs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/delvs.json")
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

	ar, err := client.DeleteVs(1)
	ok(t, err)

	equals(t, ar.Status, "ok")
	equals(t, ar.Message, "Command completed ok")
	equals(t, ar.Code, 200)

}

func TestCreateVs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/addvs.json")
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

	v := &Vs{
		Address:  "192.168.1.235",
		Protocol: "tcp",
		Port:     "6443",
	}

	vs, err := client.CreateVs(v)
	ok(t, err)

	equals(t, vs.Address, "192.168.1.235")
	equals(t, vs.Protocol, "tcp")
	equals(t, vs.VSPort, "6443")

}

func TestCreateVsIntegration(t *testing.T) {
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
		Address:  "192.168.1.112",
		Port:     "80",
		NickName: "IntegrationTest",
		Enable:   true,
		/*	  SSLReverse:    d.Get("sslreverse").(bool),
			  SSLReencrypt:  d.Get("sslreencrypt").(bool),
			  InterceptMode: d.Get("interceptmode").(int),
			  Intercept:     d.Get("intercept").(bool),
			  ForceL4:       d.Get("forcel4").(bool), */
		ForceL7:  true,
		Type:     "gen",
		Protocol: "tcp",
	}
	cvs, err := client.CreateVs(vs)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	t.Cleanup(func() {
		_, err := client.DeleteVs(cvs.Index)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
}

func TestModifyVs(t *testing.T) {
	content, err := ioutil.ReadFile("test_data/modvs.json")
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
	v := &Vs{
		Index:    1,
		Address:  "192.168.1.215",
		Protocol: "tcp",
		Port:     "6443",
	}

	vs, err := client.ModifyVs(v)
	ok(t, err)

	equals(t, vs.Address, "192.168.1.215")
	equals(t, vs.Protocol, "tcp")
	equals(t, vs.VSPort, "6453")

}

func TestModifyVsIntegration(t *testing.T) {
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
		/*	  SSLReverse:    d.Get("sslreverse").(bool),
			  SSLReencrypt:  d.Get("sslreencrypt").(bool),
			  InterceptMode: d.Get("interceptmode").(int),
			  Intercept:     d.Get("intercept").(bool),
			  ForceL4:       d.Get("forcel4").(bool), */
		ForceL7:  true,
		Type:     "gen",
		Protocol: "tcp",
	}
	cvs, err := client.CreateVs(vs)
	ok(t, err)

	v := &Vs{
		Index:    cvs.Index,
		Address:  "192.168.1.215",
		Protocol: "tcp",
		Port:     "6443",
		VSPort:   "6543",
	}

	mvs, err := client.ModifyVs(v)
	ok(t, err)

	t.Cleanup(func() {
		_, err := client.DeleteVs(mvs.Index)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
}

func TestGetVsIntegration(t *testing.T) {
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
		/*	  SSLReverse:    d.Get("sslreverse").(bool),
			  SSLReencrypt:  d.Get("sslreencrypt").(bool),
			  InterceptMode: d.Get("interceptmode").(int),
			  Intercept:     d.Get("intercept").(bool),
			  ForceL4:       d.Get("forcel4").(bool), */
		ForceL7:  true,
		Type:     "gen",
		Protocol: "tcp",
	}
	cvs, err := client.CreateVs(vs)
	ok(t, err)

	gvs, err := client.GetVs(cvs.Index)
	ok(t, err)

	t.Cleanup(func() {
		_, err := client.DeleteVs(gvs.Index)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
}

func TestMarshalJSON(t *testing.T) {
	var vs VsApiPayLoad
	vs.Address = "192.168.1.10"
	vs.Port = "888"
	vs.Protocol = "tcp"
	vs.ApiKey = "bar"
	vs.CMD = "addvs"
	vs.InterceptOpts = []string{"opnormal", "auditnone"}

	ret, err := json.Marshal(vs)
	ok(t, err)
	equals(t, string(ret), "{\"apikey\":\"bar\",\"cmd\":\"addvs\",\"vs\":\"192.168.1.10\",\"port\":\"888\",\"InterceptOpts\":\"opnormal;auditnone\",\"prot\":\"tcp\"}")
}
