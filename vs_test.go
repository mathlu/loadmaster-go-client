package lmclient

import (
	"encoding/json"
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
	equals(t, vs.Port, "6443")

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
	equals(t, vs.Port, "6453")

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
	equals(t, string(ret), "{\"apikey\":\"bar\",\"cmd\":\"addvs\",\"vs\":\"192.168.1.10\",\"port\":\"888\",\"NickName\":\"\",\"Enable\":false,\"SSLReverse\":false,\"SSLReencrypt\":false,\"InterceptMode\":0,\"Intercept\":false,\"InterceptOpts\":\"opnormal;auditnone\",\"AlertThreshold\":0,\"OwaspOpts\":\"\",\"BlockingParanoia\":0,\"IPReputationBlocking\":false,\"ExecutingParanoia\":0,\"AnomalyScoringThreshold\":0,\"PCRELimit\":0,\"JSONDLimit\":0,\"BodyLimit\":0,\"Transactionlimit\":0,\"Transparent\":false,\"SubnetOriginating\":false,\"ServerInit\":0,\"StartTLSMode\":0,\"Idletime\":0,\"Cache\":false,\"Compress\":false,\"Verify\":0,\"UseforSnat\":false,\"ForceL4\":false,\"ForceL7\":false,\"MultiConnect\":false,\"ClientCert\":0,\"SecurityHeaderOptions\":0,\"SameSite\":0,\"VerifyBearer\":false,\"ErrorCode\":\"\",\"CheckUse1.1\":false,\"MatchLen\":0,\"CheckUseGet\":0,\"SSLRewrite\":\"\",\"VStype\":\"\",\"FollowVSID\":0,\"Protocol\":\"tcp\",\"Schedule\":\"\",\"CheckType\":\"\",\"PersistTimeout\":\"\",\"CheckPort\":\"\",\"HTTPReschedule\":false,\"NRules\":0,\"NRequestRules\":0,\"NResponseRules\":0,\"NMatchBodyRules\":0,\"NPreProcessRules\":0,\"EspEnabled\":false,\"InputAuthMode\":0,\"OutputAuthMode\":0,\"MasterVS\":0,\"MasterVSID\":0,\"IsTransparent\":0,\"AddVia\":0,\"QoS\":0,\"TlsType\":\"\",\"NeedHostName\":false,\"OCSPVerify\":false,\"AllowHTTP2\":false,\"PassCipher\":false,\"PassSni\":false,\"ChkInterval\":0,\"ChkTimeout\":0,\"ChkRetryCount\":0,\"Bandwidth\":0,\"ConnsPerSecLimit\":0,\"RequestsPerSecLimit\":0,\"MaxConnsLimit\":0,\"RefreshPersist\":false,\"EnhancedHealthChecks\":false,\"RsMinimum\":0,\"NumberOfRSs\":0}")
}
