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
			client := Client{server.Client(), "bar", server.URL, tc.apiversion}

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
			client := Client{server.Client(), "bar", server.URL, tc.apiversion}

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
			client := Client{server.Client(), "bar", server.URL, tc.apiversion}

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
		{1, "/access/addvs?apikey=bar&port=6443&prot=tcp&vs=192.168.1.235", "test_data/addvs.xml"},
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
			client := Client{server.Client(), "bar", server.URL, tc.apiversion}

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
			client := NewClient(os.Getenv("LOADMASTER_API_KEY"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")), tc.apiversion)
			if vi, err := client.GetVsByName("IntegrationTestV" + strconv.Itoa(tc.apiversion)); err == nil {
				_, err := client.DeleteVs(vi.Index)
				if err != nil {
					fmt.Printf("err: %v", err)
				}
			}
			vs := &Vs{
				Address:  "192.168.1.112",
				Port:     "80",
				NickName: "IntegrationTest" + strconv.Itoa(tc.apiversion),
				Type:     "gen",
				Protocol: "tcp",
			}
			vsc, err := client.CreateVs(vs)
			index := vsc.Index
			if err != nil {
				fmt.Printf("err: %v", err)
			}
			rss := []Rs{
				Rs{
					Addr:    "192.168.1.50",
					Port:    80,
					VSIndex: index,
				},
				Rs{
					Addr:    "192.168.1.55",
					Port:    80,
					VSIndex: index,
				},
			}
			for _, rs := range rss {
				_, err = client.CreateRs(&rs)
				if err != nil {
					fmt.Printf("err: %v", err)
				}
			}
			t.Cleanup(func() {
				_, err := client.DeleteVs(index)
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
		{1, "/access/modvs?apikey=bar&port=6443&prot=tcp&vs=1&vsaddress=192.168.1.215", "test_data/modvs.xml"},
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
			client := Client{server.Client(), "bar", server.URL, tc.apiversion}
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
			client := NewClient(os.Getenv("LOADMASTER_API_KEY"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")), tc.apiversion)
			if vi, err := client.GetVsByName("IntegrationTestV" + strconv.Itoa(tc.apiversion)); err == nil {
				_, err := client.DeleteVs(vi.Index)
				ok(t, err)
			}
			vs := &Vs{
				Address:  "192.168.1.112",
				Port:     "80",
				NickName: "IntegrationTest" + strconv.Itoa(tc.apiversion),
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
			client := NewClient(os.Getenv("LOADMASTER_API_KEY"), fmt.Sprintf("https://%s/", os.Getenv("LOADMASTER_SERVER")), tc.apiversion)
			if vi, err := client.GetVsByName("IntegrationTestV" + strconv.Itoa(tc.apiversion)); err == nil {
				_, err := client.DeleteVs(vi.Index)
				ok(t, err)
			}
			vs := &Vs{
				Address:  "192.168.1.112",
				Port:     "80",
				NickName: "IntegrationTest" + strconv.Itoa(tc.apiversion),
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
