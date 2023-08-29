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

func TestGetAllVs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/listvs.json"},
		{1, "/access/listvs?apikey=bar", "test_data/listvs.xml"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			content, err := ioutil.ReadFile(tc.datafile)
			ok(t, err)
			// Start a local HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				equals(t, req.URL.String(), tc.url)
				// Send response to be tested
				_, err := rw.Write([]byte(content))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}
			}))

			defer server.Close()
			client := Client{server.Client(), "bar", "foo", "baz", server.URL, tc.apiversion}

			vs, err := client.GetAllVs()
			ok(t, err)

			equals(t, vs[0].Index, 1)
			equals(t, vs[0].NickName, "foo")
		})
	}

}

func TestGetVs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/showvs.json"},
		{1, "/access/showvs?apikey=bar&vs=1", "test_data/showvs.xml"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			content, err := ioutil.ReadFile(tc.datafile)
			ok(t, err)
			// Start a local HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				equals(t, req.URL.String(), tc.url)
				// Send response to be tested
				_, err := rw.Write([]byte(content))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}
			}))

			defer server.Close()
			client := Client{server.Client(), "bar", "foo", "baz", server.URL, tc.apiversion}

			vs, err := client.GetVs(1)
			ok(t, err)

			equals(t, vs.Index, 1)
			equals(t, vs.NickName, "foo")
		})
	}

}

func TestDelVs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/delvs.json"},
		{1, "/access/delvs?apikey=bar&vs=1", "test_data/delvs.xml"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			content, err := ioutil.ReadFile(tc.datafile)
			ok(t, err)
			// Start a local HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				equals(t, req.URL.String(), tc.url)
				// Send response to be tested
				_, err := rw.Write([]byte(content))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}
			}))

			defer server.Close()
			client := Client{server.Client(), "bar", "foo", "baz", server.URL, tc.apiversion}

			ar, err := client.DeleteVs(1)
			ok(t, err)

			equals(t, ar.Status, "ok")
			equals(t, ar.Message, "Command completed ok")
			equals(t, ar.Code, 200)
		})
	}

}

func TestCreateVs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/addvs.json"},
		{1, "/access/addvs?Enable=Y&apikey=bar&defaultgw=192.168.1.1&forcel4=1&forcel7=0&port=6443&prot=tcp&vs=192.168.1.235&vstype=http2", "test_data/addvs.xml"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			content, err := ioutil.ReadFile(tc.datafile)
			ok(t, err)
			// Start a local HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				equals(t, req.URL.String(), tc.url)
				// Send response to be tested
				_, err := rw.Write([]byte(content))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}
			}))

			defer server.Close()
			client := Client{server.Client(), "bar", "foo", "baz", server.URL, tc.apiversion}

			v := &Vs{
				Address:   "192.168.1.235",
				Protocol:  "tcp",
				Port:      "6443",
				Type:      "http2",
				Enable:    true,
				Layer:     4,
				DefaultGW: "192.168.1.1",
			}

			vs, err := client.CreateVs(v)
			ok(t, err)

			equals(t, vs.Address, "192.168.1.235")
			equals(t, vs.Protocol, "tcp")
			equals(t, vs.VSPort, "6443")
			equals(t, vs.Type, "http2")
			equals(t, vs.Enable, true)
			equals(t, vs.ForceL4, true)
			equals(t, vs.ForceL7, false)
			equals(t, vs.DefaultGW, "192.168.1.1")
			equals(t, vs.Layer, 4)
		})
	}

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
	testCases := []struct {
		apiversion int
	}{
		{2},
		{1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			client := NewClient(os.Getenv("LOADMASTER_API_KEY"), os.Getenv("LOADMASTER_API_USER"), os.Getenv("LOADMASTER_API_PASS"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")), tc.apiversion)
			if vi, err := client.GetVsByName("IntegrationTestV" + strconv.Itoa(tc.apiversion)); err == nil {
				_, err := client.DeleteVs(vi.Index)
				if err != nil {
					fmt.Printf("err: %v", err)
				}
			}
			vs := &Vs{
				Address:   "192.168.1.112",
				Port:      "80",
				NickName:  "IntegrationTestV" + strconv.Itoa(tc.apiversion),
				Type:      "gen",
				Protocol:  "tcp",
				Enable:    true,
				Layer:     4,
				DefaultGW: "192.168.1.1",
			}
			vsc, err := client.CreateVs(vs)
			if err != nil {
				fmt.Printf("err: %v", err)
			}

			equals(t, vsc.Address, "192.168.1.112")
			equals(t, vsc.Protocol, "tcp")
			equals(t, vsc.VSPort, "80")
			equals(t, vsc.Type, "gen")
			equals(t, vsc.Enable, true)
			equals(t, vsc.ForceL4, true)
			equals(t, vsc.ForceL7, false)
			equals(t, vsc.Layer, 4)
			equals(t, vsc.DefaultGW, "192.168.1.1")

			t.Cleanup(func() {
				_, err := client.DeleteVs(vsc.Index)
				if err != nil {
					fmt.Printf("err: %v", err)
				}
			})
		})
	}
}

func TestModifyVs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/modvs.json"},
		{1, "/access/modvs?Enable=Y&apikey=bar&defaultgw=192.168.1.1&forcel4=0&forcel7=1&port=6443&prot=tcp&vs=1&vsaddress=192.168.1.215", "test_data/modvs.xml"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			content, err := ioutil.ReadFile(tc.datafile)
			ok(t, err)
			// Start a local HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				equals(t, req.URL.String(), tc.url)
				// Send response to be tested
				_, err := rw.Write([]byte(content))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}
			}))

			defer server.Close()
			client := Client{server.Client(), "bar", "foo", "baz", server.URL, tc.apiversion}
			v := &Vs{
				Index:     1,
				Address:   "192.168.1.215",
				Protocol:  "tcp",
				Port:      "6443",
				Enable:    true,
				Layer:     7,
				DefaultGW: "192.168.1.1",
			}

			vs, err := client.ModifyVs(v)
			ok(t, err)

			equals(t, vs.Address, "192.168.1.215")
			equals(t, vs.Protocol, "tcp")
			equals(t, vs.VSPort, "6453")
			equals(t, vs.Type, "gen")
			equals(t, vs.Enable, true)
			equals(t, vs.ForceL4, false)
			equals(t, vs.ForceL7, true)
			equals(t, vs.DefaultGW, "192.168.1.1")
			equals(t, vs.Layer, 7)
		})
	}

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
	testCases := []struct {
		apiversion int
	}{
		{2},
		{1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			client := NewClient(os.Getenv("LOADMASTER_API_KEY"), os.Getenv("LOADMASTER_API_USER"), os.Getenv("LOADMASTER_API_PASS"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")), tc.apiversion)
			if vi, err := client.GetVsByName("IntegrationTestV" + strconv.Itoa(tc.apiversion)); err == nil {
				_, err := client.DeleteVs(vi.Index)
				ok(t, err)
			}
			vs := &Vs{
				Address:  "192.168.1.112",
				Port:     "80",
				NickName: "IntegrationTestV" + strconv.Itoa(tc.apiversion),
				Type:     "gen",
				Protocol: "tcp",
				Enable:   true,
				Layer:    7,
			}
			cvs, err := client.CreateVs(vs)
			ok(t, err)

			equals(t, cvs.Address, "192.168.1.112")
			equals(t, cvs.Protocol, "tcp")
			equals(t, cvs.VSPort, "80")
			equals(t, cvs.Type, "gen")
			equals(t, cvs.Enable, true)
			equals(t, cvs.ForceL4, false)
			equals(t, cvs.ForceL7, true)
			equals(t, cvs.Layer, 7)

			v := &Vs{
				Index:    cvs.Index,
				Address:  "192.168.1.215",
				Protocol: "tcp",
				Port:     "6443",
				VSPort:   "6543",
				Enable:   true,
			}

			mvs, err := client.ModifyVs(v)
			ok(t, err)

			equals(t, mvs.Address, "192.168.1.215")
			equals(t, mvs.Protocol, "tcp")
			equals(t, mvs.VSPort, "6543")
			equals(t, mvs.Type, "gen")
			equals(t, mvs.Enable, true)
			equals(t, mvs.ForceL4, false)
			equals(t, mvs.ForceL7, true)
			equals(t, mvs.Layer, 7)

			t.Cleanup(func() {
				_, err := client.DeleteVs(mvs.Index)
				if err != nil {
					fmt.Printf("err: %v", err)
				}
			})
		})
	}
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
	testCases := []struct {
		apiversion int
	}{
		{2},
		{1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("apiversion_%d", tc.apiversion), func(t *testing.T) {
			client := NewClient(os.Getenv("LOADMASTER_API_KEY"), os.Getenv("LOADMASTER_API_USER"), os.Getenv("LOADMASTER_API_PASS"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")), tc.apiversion)
			if vi, err := client.GetVsByName("IntegrationTestV" + strconv.Itoa(tc.apiversion)); err == nil {
				_, err := client.DeleteVs(vi.Index)
				ok(t, err)
			}
			vs := &Vs{
				Address:  "192.168.1.112",
				Port:     "80",
				NickName: "IntegrationTestV" + strconv.Itoa(tc.apiversion),
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
		})
	}
}
