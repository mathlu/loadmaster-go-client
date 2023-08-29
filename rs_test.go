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
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/showrs.json"},
		{1, "/access/showrs?apikey=bar&rs=%211&vs=1", "test_data/showrs.xml"},
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

			rs, err := client.GetRs(1, 1)
			ok(t, err)

			equals(t, rs.RsIndex, 1)
			equals(t, rs.Addr, "10.10.10.10")
		})
	}

}

func TestDelRs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/delrs.json"},
		{1, "/access/delrs?apikey=bar&rs=%211&vs=1", "test_data/delrs.xml"},
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

			ar, err := client.DeleteRs(1, 1)
			ok(t, err)

			equals(t, ar.Status, "ok")
			equals(t, ar.Message, "Command completed ok")
			equals(t, ar.Code, 200)
		})
	}

}

func TestCreateRs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/addrs.json"},
		{1, "/access/addrs?apikey=bar&non_local=1&rs=10.10.10.10&rsport=8080&vs=1", "test_data/addrs.xml"},
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

			r := &Rs{
				VSIndex: 1,
				Addr:    "10.10.10.10",
				Port:    8080,
			}

			rs, err := client.CreateRs(r)
			ok(t, err)

			equals(t, rs.Addr, "10.10.10.10")
			equals(t, rs.Port, 8080)
		})
	}

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
				Address:  "192.168.1.124",
				Port:     "8093",
				NickName: "IntegrationTest" + strconv.Itoa(tc.apiversion),
				Type:     "gen",
				Protocol: "tcp",
				Enable:   true,
			}
			vsc, err := client.CreateVs(vs)
			if err != nil {
				fmt.Printf("err: %v", err)
			}
			index := vsc.Index
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
		})
	}
}

func TestModifyRs(t *testing.T) {
	testCases := []struct {
		apiversion int
		url        string
		datafile   string
	}{
		{2, "/accessv2", "test_data/modrs.json"},
		{1, "/access/modrs?apikey=bar&newport=6443&non_local=1&rs=%211&vs=1", "test_data/modrs.xml"},
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
			r := &Rs{
				VSIndex: 1,
				RsIndex: 1,
				Addr:    "192.168.1.215",
				NewPort: "6443",
			}

			ar, err := client.ModifyRs(r)
			ok(t, err)

			equals(t, ar.Status, "ok")
			equals(t, ar.Message, "Command completed ok")
			equals(t, ar.Code, 200)
		})
	}
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
				Address:  "192.168.1.123",
				Port:     "8090",
				NickName: "IntegrationTestV" + strconv.Itoa(tc.apiversion),
				Type:     "gen",
				Protocol: "tcp",
				Enable:   true,
			}
			cvs, err := client.CreateVs(vs)
			ok(t, err)
			if err != nil {
				fmt.Printf("err: %v", err)
			}
			index := cvs.Index
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
				RsIndex: rsc.RsIndex,
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
		})
	}
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
		})
	}
}
