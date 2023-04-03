package lmclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Vss struct {
	VS []Vs
}
type Vs struct {
	Status                  string   `json:"Status"`
	Index                   int      `json:"Index"`
	Address                 string   `json:"VSAddress"`
	Port                    string   `json:"VSPort"`
	Layer                   int      `json:"Layer"`
	NickName                string   `json:"NickName"`
	Enable                  bool     `json:"Enable"`
	SSLReverse              bool     `json:"SSLReverse"`
	SSLReencrypt            bool     `json:"SSLReencrypt"`
	InterceptMode           int      `json:"InterceptMode"`
	Intercept               bool     `json:"Intercept"`
	InterceptOpts           []string `json:"InterceptOpts"`
	AlertThreshold          int      `json:"AlertThreshold"`
	OwaspOpts               []string `json:"OwaspOpts"`
	BlockingParanoia        int      `json:"BlockingParanoia"`
	IPReputationBlocking    bool     `json:"IPReputationBlocking"`
	ExecutingParanoia       int      `json:"ExecutingParanoia"`
	AnomalyScoringThreshold int      `json:"AnomalyScoringThreshold"`
	PCRELimit               int      `json:"PCRELimit"`
	JSONDLimit              int      `json:"JSONDLimit"`
	BodyLimit               int      `json:"BodyLimit"`
	Transactionlimit        int      `json:"Transactionlimit"`
	Transparent             bool     `json:"Transparent"`
	SubnetOriginating       bool     `json:"SubnetOriginating"`
	ServerInit              int      `json:"ServerInit"`
	StartTLSMode            int      `json:"StartTLSMode"`
	Idletime                int      `json:"Idletime"`
	Cache                   bool     `json:"Cache"`
	Compress                bool     `json:"Compress"`
	Verify                  int      `json:"Verify"`
	UseforSnat              bool     `json:"UseforSnat"`
	ForceL4                 bool     `json:"ForceL4"`
	ForceL7                 bool     `json:"ForceL7"`
	MultiConnect            bool     `json:"MultiConnect"`
	ClientCert              int      `json:"ClientCert"`
	SecurityHeaderOptions   int      `json:"SecurityHeaderOptions"`
	SameSite                int      `json:"SameSite"`
	VerifyBearer            bool     `json:"VerifyBearer"`
	ErrorCode               string   `json:"ErrorCode"`
	CheckUse11              bool     `json:"CheckUse1.1"`
	MatchLen                int      `json:"MatchLen"`
	CheckUseGet             int      `json:"CheckUseGet"`
	SSLRewrite              string   `json:"SSLRewrite"`
	Type                    string   `json:"VStype"`
	FollowVSID              int      `json:"FollowVSID"`
	Protocol                string   `json:"Protocol"`
	Schedule                string   `json:"Schedule"`
	CheckType               string   `json:"CheckType"`
	PersistTimeout          string   `json:"PersistTimeout"`
	CheckPort               string   `json:"CheckPort"`
	HTTPReschedule          bool     `json:"HTTPReschedule"`
	NRules                  int      `json:"NRules"`
	NRequestRules           int      `json:"NRequestRules"`
	NResponseRules          int      `json:"NResponseRules"`
	NMatchBodyRules         int      `json:"NMatchBodyRules"`
	NPreProcessRules        int      `json:"NPreProcessRules"`
	EspEnabled              bool     `json:"EspEnabled"`
	InputAuthMode           int      `json:"InputAuthMode"`
	OutputAuthMode          int      `json:"OutputAuthMode"`
	MasterVS                int      `json:"MasterVS"`
	MasterVSID              int      `json:"MasterVSID"`
	IsTransparent           int      `json:"IsTransparent"`
	AddVia                  int      `json:"AddVia"`
	QoS                     int      `json:"QoS"`
	TLSType                 string   `json:"TlsType"`
	NeedHostName            bool     `json:"NeedHostName"`
	OCSPVerify              bool     `json:"OCSPVerify"`
	AllowHTTP2              bool     `json:"AllowHTTP2"`
	PassCipher              bool     `json:"PassCipher"`
	PassSni                 bool     `json:"PassSni"`
	ChkInterval             int      `json:"ChkInterval"`
	ChkTimeout              int      `json:"ChkTimeout"`
	ChkRetryCount           int      `json:"ChkRetryCount"`
	Bandwidth               int      `json:"Bandwidth"`
	ConnsPerSecLimit        int      `json:"ConnsPerSecLimit"`
	RequestsPerSecLimit     int      `json:"RequestsPerSecLimit"`
	MaxConnsLimit           int      `json:"MaxConnsLimit"`
	RefreshPersist          bool     `json:"RefreshPersist"`
	EnhancedHealthChecks    bool     `json:"EnhancedHealthChecks"`
	RsMinimum               int      `json:"RsMinimum"`
	NumberOfRSs             int      `json:"NumberOfRSs"`
}

type ApiPayLoad struct {
	ApiKey string `json:"apikey"`
	CMD    string `json:"cmd"`
}

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type VsApiPayLoad struct {
	Vs
	ApiPayLoad
}

func (u VsApiPayLoad) MarshalJSON() ([]byte, error) {
	switch u.CMD {
	case "addvs":
		var interceptopts string
		var owaspopts string
		if len(u.InterceptOpts) > 0 {
			interceptopts = strings.Join(u.InterceptOpts, ";")
		}
		if len(u.OwaspOpts) > 0 {
			owaspopts = strings.Join(u.OwaspOpts, ";")
		}
		return json.Marshal(&struct {
			ApiKey                  string `json:"apikey"`
			CMD                     string `json:"cmd"`
			Address                 string `json:"vs"`
			Port                    string `json:"port"`
			NickName                string `json:"NickName"`
			Enable                  bool   `json:"Enable"`
			SSLReverse              bool   `json:"SSLReverse"`
			SSLReencrypt            bool   `json:"SSLReencrypt"`
			InterceptMode           int    `json:"InterceptMode"`
			Intercept               bool   `json:"Intercept"`
			InterceptOpts           string `json:"InterceptOpts"`
			AlertThreshold          int    `json:"AlertThreshold"`
			OwaspOpts               string `json:"OwaspOpts"`
			BlockingParanoia        int    `json:"BlockingParanoia"`
			IPReputationBlocking    bool   `json:"IPReputationBlocking"`
			ExecutingParanoia       int    `json:"ExecutingParanoia"`
			AnomalyScoringThreshold int    `json:"AnomalyScoringThreshold"`
			PCRELimit               int    `json:"PCRELimit"`
			JSONDLimit              int    `json:"JSONDLimit"`
			BodyLimit               int    `json:"BodyLimit"`
			Transactionlimit        int    `json:"Transactionlimit"`
			Transparent             bool   `json:"Transparent"`
			SubnetOriginating       bool   `json:"SubnetOriginating"`
			ServerInit              int    `json:"ServerInit"`
			StartTLSMode            int    `json:"StartTLSMode"`
			Idletime                int    `json:"Idletime"`
			Cache                   bool   `json:"Cache"`
			Compress                bool   `json:"Compress"`
			Verify                  int    `json:"Verify"`
			UseforSnat              bool   `json:"UseforSnat"`
			ForceL4                 bool   `json:"ForceL4"`
			ForceL7                 bool   `json:"ForceL7"`
			MultiConnect            bool   `json:"MultiConnect"`
			ClientCert              int    `json:"ClientCert"`
			SecurityHeaderOptions   int    `json:"SecurityHeaderOptions"`
			SameSite                int    `json:"SameSite"`
			VerifyBearer            bool   `json:"VerifyBearer"`
			ErrorCode               string `json:"ErrorCode"`
			CheckUse11              bool   `json:"CheckUse1.1"`
			MatchLen                int    `json:"MatchLen"`
			CheckUseGet             int    `json:"CheckUseGet"`
			SSLRewrite              string `json:"SSLRewrite"`
			Type                    string `json:"VStype"`
			FollowVSID              int    `json:"FollowVSID"`
			Protocol                string `json:"Protocol"`
			Schedule                string `json:"Schedule"`
			CheckType               string `json:"CheckType"`
			PersistTimeout          string `json:"PersistTimeout"`
			CheckPort               string `json:"CheckPort"`
			HTTPReschedule          bool   `json:"HTTPReschedule"`
			NRules                  int    `json:"NRules"`
			NRequestRules           int    `json:"NRequestRules"`
			NResponseRules          int    `json:"NResponseRules"`
			NMatchBodyRules         int    `json:"NMatchBodyRules"`
			NPreProcessRules        int    `json:"NPreProcessRules"`
			EspEnabled              bool   `json:"EspEnabled"`
			InputAuthMode           int    `json:"InputAuthMode"`
			OutputAuthMode          int    `json:"OutputAuthMode"`
			MasterVS                int    `json:"MasterVS"`
			MasterVSID              int    `json:"MasterVSID"`
			IsTransparent           int    `json:"IsTransparent"`
			AddVia                  int    `json:"AddVia"`
			QoS                     int    `json:"QoS"`
			TLSType                 string `json:"TlsType"`
			NeedHostName            bool   `json:"NeedHostName"`
			OCSPVerify              bool   `json:"OCSPVerify"`
			AllowHTTP2              bool   `json:"AllowHTTP2"`
			PassCipher              bool   `json:"PassCipher"`
			PassSni                 bool   `json:"PassSni"`
			ChkInterval             int    `json:"ChkInterval"`
			ChkTimeout              int    `json:"ChkTimeout"`
			ChkRetryCount           int    `json:"ChkRetryCount"`
			Bandwidth               int    `json:"Bandwidth"`
			ConnsPerSecLimit        int    `json:"ConnsPerSecLimit"`
			RequestsPerSecLimit     int    `json:"RequestsPerSecLimit"`
			MaxConnsLimit           int    `json:"MaxConnsLimit"`
			RefreshPersist          bool   `json:"RefreshPersist"`
			EnhancedHealthChecks    bool   `json:"EnhancedHealthChecks"`
			RsMinimum               int    `json:"RsMinimum"`
			NumberOfRSs             int    `json:"NumberOfRSs"`
		}{
			ApiKey:                  u.ApiKey,
			CMD:                     u.CMD,
			Address:                 u.Address,
			Port:                    u.Port,
			NickName:                u.NickName,
			SSLReverse:              u.SSLReverse,
			SSLReencrypt:            u.SSLReencrypt,
			InterceptMode:           u.InterceptMode,
			Intercept:               u.Intercept,
			InterceptOpts:           interceptopts,
			AlertThreshold:          u.AlertThreshold,
			OwaspOpts:               owaspopts,
			BlockingParanoia:        u.BlockingParanoia,
			IPReputationBlocking:    u.IPReputationBlocking,
			ExecutingParanoia:       u.ExecutingParanoia,
			AnomalyScoringThreshold: u.AnomalyScoringThreshold,
			PCRELimit:               u.PCRELimit,
			JSONDLimit:              u.JSONDLimit,
			BodyLimit:               u.BodyLimit,
			Transactionlimit:        u.Transactionlimit,
			Transparent:             u.Transparent,
			SubnetOriginating:       u.SubnetOriginating,
			ServerInit:              u.ServerInit,
			StartTLSMode:            u.StartTLSMode,
			Idletime:                u.Idletime,
			Cache:                   u.Cache,
			Compress:                u.Compress,
			Verify:                  u.Verify,
			UseforSnat:              u.UseforSnat,
			ForceL4:                 u.ForceL4,
			ForceL7:                 u.ForceL7,
			MultiConnect:            u.MultiConnect,
			ClientCert:              u.ClientCert,
			SecurityHeaderOptions:   u.SecurityHeaderOptions,
			SameSite:                u.SameSite,
			VerifyBearer:            u.VerifyBearer,
			ErrorCode:               u.ErrorCode,
			CheckUse11:              u.CheckUse11,
			MatchLen:                u.MatchLen,
			CheckUseGet:             u.CheckUseGet,
			SSLRewrite:              u.SSLRewrite,
			Type:                    u.Type,
			FollowVSID:              u.FollowVSID,
			Protocol:                u.Protocol,
			Schedule:                u.Schedule,
			CheckType:               u.CheckType,
			PersistTimeout:          u.PersistTimeout,
			CheckPort:               u.CheckPort,
			HTTPReschedule:          u.HTTPReschedule,
			NRules:                  u.NRules,
			NRequestRules:           u.NRequestRules,
			NResponseRules:          u.NResponseRules,
			NMatchBodyRules:         u.NMatchBodyRules,
			NPreProcessRules:        u.NPreProcessRules,
			EspEnabled:              u.EspEnabled,
			InputAuthMode:           u.InputAuthMode,
			OutputAuthMode:          u.OutputAuthMode,
			MasterVS:                u.MasterVS,
			MasterVSID:              u.MasterVSID,
			IsTransparent:           u.IsTransparent,
			AddVia:                  u.AddVia,
			QoS:                     u.QoS,
			TLSType:                 u.TLSType,
			NeedHostName:            u.NeedHostName,
			OCSPVerify:              u.OCSPVerify,
			AllowHTTP2:              u.AllowHTTP2,
			PassCipher:              u.PassCipher,
			PassSni:                 u.PassSni,
			ChkInterval:             u.ChkInterval,
			ChkTimeout:              u.ChkTimeout,
			ChkRetryCount:           u.ChkRetryCount,
			Bandwidth:               u.Bandwidth,
			ConnsPerSecLimit:        u.ConnsPerSecLimit,
			RequestsPerSecLimit:     u.RequestsPerSecLimit,
			MaxConnsLimit:           u.MaxConnsLimit,
			RefreshPersist:          u.RefreshPersist,
			EnhancedHealthChecks:    u.EnhancedHealthChecks,
			RsMinimum:               u.RsMinimum,
			NumberOfRSs:             u.NumberOfRSs,
		})
	case "delvs", "showvs":
		return json.Marshal(&struct {
			Index  int    `json:"vs"`
			ApiKey string `json:"apikey"`
			CMD    string `json:"cmd"`
		}{
			Index:  u.Index,
			ApiKey: u.ApiKey,
			CMD:    u.CMD,
		})
	case "modvs":
		var interceptopts string
		var owaspopts string
		if len(u.InterceptOpts) > 0 {
			interceptopts = strings.Join(u.InterceptOpts, ";")
		}
		if len(u.OwaspOpts) > 0 {
			owaspopts = strings.Join(u.OwaspOpts, ";")
		}
		return json.Marshal(&struct {
			ApiKey                  string `json:"apikey"`
			CMD                     string `json:"cmd"`
			Index                   int    `json:"vs"`
			Address                 string `json:"vsaddress"`
			Port                    string `json:"vsport"`
			NickName                string `json:"NickName"`
			Enable                  bool   `json:"Enable"`
			SSLReverse              bool   `json:"SSLReverse"`
			SSLReencrypt            bool   `json:"SSLReencrypt"`
			InterceptMode           int    `json:"InterceptMode"`
			Intercept               bool   `json:"Intercept"`
			InterceptOpts           string `json:"InterceptOpts"`
			AlertThreshold          int    `json:"AlertThreshold"`
			OwaspOpts               string `json:"OwaspOpts"`
			BlockingParanoia        int    `json:"BlockingParanoia"`
			IPReputationBlocking    bool   `json:"IPReputationBlocking"`
			ExecutingParanoia       int    `json:"ExecutingParanoia"`
			AnomalyScoringThreshold int    `json:"AnomalyScoringThreshold"`
			PCRELimit               int    `json:"PCRELimit"`
			JSONDLimit              int    `json:"JSONDLimit"`
			BodyLimit               int    `json:"BodyLimit"`
			Transactionlimit        int    `json:"Transactionlimit"`
			Transparent             bool   `json:"Transparent"`
			SubnetOriginating       bool   `json:"SubnetOriginating"`
			ServerInit              int    `json:"ServerInit"`
			StartTLSMode            int    `json:"StartTLSMode"`
			Idletime                int    `json:"Idletime"`
			Cache                   bool   `json:"Cache"`
			Compress                bool   `json:"Compress"`
			Verify                  int    `json:"Verify"`
			UseforSnat              bool   `json:"UseforSnat"`
			ForceL4                 bool   `json:"ForceL4"`
			ForceL7                 bool   `json:"ForceL7"`
			MultiConnect            bool   `json:"MultiConnect"`
			ClientCert              int    `json:"ClientCert"`
			SecurityHeaderOptions   int    `json:"SecurityHeaderOptions"`
			SameSite                int    `json:"SameSite"`
			VerifyBearer            bool   `json:"VerifyBearer"`
			ErrorCode               string `json:"ErrorCode"`
			CheckUse11              bool   `json:"CheckUse1.1"`
			MatchLen                int    `json:"MatchLen"`
			CheckUseGet             int    `json:"CheckUseGet"`
			SSLRewrite              string `json:"SSLRewrite"`
			Type                    string `json:"VStype"`
			FollowVSID              int    `json:"FollowVSID"`
			Protocol                string `json:"Protocol"`
			Schedule                string `json:"Schedule"`
			CheckType               string `json:"CheckType"`
			PersistTimeout          string `json:"PersistTimeout"`
			CheckPort               string `json:"CheckPort"`
			HTTPReschedule          bool   `json:"HTTPReschedule"`
			NRules                  int    `json:"NRules"`
			NRequestRules           int    `json:"NRequestRules"`
			NResponseRules          int    `json:"NResponseRules"`
			NMatchBodyRules         int    `json:"NMatchBodyRules"`
			NPreProcessRules        int    `json:"NPreProcessRules"`
			EspEnabled              bool   `json:"EspEnabled"`
			InputAuthMode           int    `json:"InputAuthMode"`
			OutputAuthMode          int    `json:"OutputAuthMode"`
			MasterVS                int    `json:"MasterVS"`
			MasterVSID              int    `json:"MasterVSID"`
			IsTransparent           int    `json:"IsTransparent"`
			AddVia                  int    `json:"AddVia"`
			QoS                     int    `json:"QoS"`
			TLSType                 string `json:"TlsType"`
			NeedHostName            bool   `json:"NeedHostName"`
			OCSPVerify              bool   `json:"OCSPVerify"`
			AllowHTTP2              bool   `json:"AllowHTTP2"`
			PassCipher              bool   `json:"PassCipher"`
			PassSni                 bool   `json:"PassSni"`
			ChkInterval             int    `json:"ChkInterval"`
			ChkTimeout              int    `json:"ChkTimeout"`
			ChkRetryCount           int    `json:"ChkRetryCount"`
			Bandwidth               int    `json:"Bandwidth"`
			ConnsPerSecLimit        int    `json:"ConnsPerSecLimit"`
			RequestsPerSecLimit     int    `json:"RequestsPerSecLimit"`
			MaxConnsLimit           int    `json:"MaxConnsLimit"`
			RefreshPersist          bool   `json:"RefreshPersist"`
			EnhancedHealthChecks    bool   `json:"EnhancedHealthChecks"`
			RsMinimum               int    `json:"RsMinimum"`
			NumberOfRSs             int    `json:"NumberOfRSs"`
		}{
			ApiKey:                  u.ApiKey,
			CMD:                     u.CMD,
			Address:                 u.Address,
			Port:                    u.Port,
			NickName:                u.NickName,
			SSLReverse:              u.SSLReverse,
			SSLReencrypt:            u.SSLReencrypt,
			InterceptMode:           u.InterceptMode,
			Intercept:               u.Intercept,
			InterceptOpts:           interceptopts,
			AlertThreshold:          u.AlertThreshold,
			OwaspOpts:               owaspopts,
			BlockingParanoia:        u.BlockingParanoia,
			IPReputationBlocking:    u.IPReputationBlocking,
			ExecutingParanoia:       u.ExecutingParanoia,
			AnomalyScoringThreshold: u.AnomalyScoringThreshold,
			PCRELimit:               u.PCRELimit,
			JSONDLimit:              u.JSONDLimit,
			BodyLimit:               u.BodyLimit,
			Transactionlimit:        u.Transactionlimit,
			Transparent:             u.Transparent,
			SubnetOriginating:       u.SubnetOriginating,
			ServerInit:              u.ServerInit,
			StartTLSMode:            u.StartTLSMode,
			Idletime:                u.Idletime,
			Cache:                   u.Cache,
			Compress:                u.Compress,
			Verify:                  u.Verify,
			UseforSnat:              u.UseforSnat,
			ForceL4:                 u.ForceL4,
			ForceL7:                 u.ForceL7,
			MultiConnect:            u.MultiConnect,
			ClientCert:              u.ClientCert,
			SecurityHeaderOptions:   u.SecurityHeaderOptions,
			SameSite:                u.SameSite,
			VerifyBearer:            u.VerifyBearer,
			ErrorCode:               u.ErrorCode,
			CheckUse11:              u.CheckUse11,
			MatchLen:                u.MatchLen,
			CheckUseGet:             u.CheckUseGet,
			SSLRewrite:              u.SSLRewrite,
			Type:                    u.Type,
			FollowVSID:              u.FollowVSID,
			Protocol:                u.Protocol,
			Schedule:                u.Schedule,
			CheckType:               u.CheckType,
			PersistTimeout:          u.PersistTimeout,
			CheckPort:               u.CheckPort,
			HTTPReschedule:          u.HTTPReschedule,
			NRules:                  u.NRules,
			NRequestRules:           u.NRequestRules,
			NResponseRules:          u.NResponseRules,
			NMatchBodyRules:         u.NMatchBodyRules,
			NPreProcessRules:        u.NPreProcessRules,
			EspEnabled:              u.EspEnabled,
			InputAuthMode:           u.InputAuthMode,
			OutputAuthMode:          u.OutputAuthMode,
			MasterVS:                u.MasterVS,
			MasterVSID:              u.MasterVSID,
			IsTransparent:           u.IsTransparent,
			AddVia:                  u.AddVia,
			QoS:                     u.QoS,
			TLSType:                 u.TLSType,
			NeedHostName:            u.NeedHostName,
			OCSPVerify:              u.OCSPVerify,
			AllowHTTP2:              u.AllowHTTP2,
			PassCipher:              u.PassCipher,
			PassSni:                 u.PassSni,
			ChkInterval:             u.ChkInterval,
			ChkTimeout:              u.ChkTimeout,
			ChkRetryCount:           u.ChkRetryCount,
			Bandwidth:               u.Bandwidth,
			ConnsPerSecLimit:        u.ConnsPerSecLimit,
			RequestsPerSecLimit:     u.RequestsPerSecLimit,
			MaxConnsLimit:           u.MaxConnsLimit,
			RefreshPersist:          u.RefreshPersist,
			EnhancedHealthChecks:    u.EnhancedHealthChecks,
			RsMinimum:               u.RsMinimum,
			NumberOfRSs:             u.NumberOfRSs,
		})
	default:
		return nil, errors.New("Unknown CMD")
	}
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

func (c *Client) GetVs(index int) (*Vs, error) {
	vsa := &VsApiPayLoad{
		Vs{
			Index: index,
		},
		ApiPayLoad{
			CMD:    "showvs",
			ApiKey: c.ApiKey,
		},
	}
	b, err := json.Marshal(vsa)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vs Vs
	err = json.Unmarshal(resp, &vs)
	if err != nil {
		return nil, err
	}

	return &vs, nil
}

func (c *Client) CreateVs(v *Vs) (*Vs, error) {
	vsa := &VsApiPayLoad{
		Vs{
			Address:       v.Address,
			Port:          v.Port,
			NickName:      v.NickName,
			SSLReverse:    v.SSLReverse,
			SSLReencrypt:  v.SSLReencrypt,
			InterceptMode: v.InterceptMode,
			Intercept:     v.Intercept,
			InterceptOpts: v.InterceptOpts,
			Enable:        v.Enable,
			ForceL4:       v.ForceL4,
			ForceL7:       v.ForceL7,
			Type:          v.Type,
			Protocol:      v.Protocol,
		},
		ApiPayLoad{
			ApiKey: c.ApiKey,
			CMD:    "addvs",
		},
	}

	b, err := json.Marshal(vsa)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vs Vs
	err = json.Unmarshal(resp, &vs)
	if err != nil {
		return nil, err
	}

	return &vs, nil
}

func (c *Client) DeleteVs(index int) (*ApiResponse, error) {
	vsa := &VsApiPayLoad{
		Vs{
			Index: index,
		},
		ApiPayLoad{
			CMD:    "delvs",
			ApiKey: c.ApiKey,
		},
	}
	b, err := json.Marshal(vsa)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var ar ApiResponse
	err = json.Unmarshal(resp, &ar)
	if err != nil {
		return nil, err
	}

	if ar.Status != "ok" {
		return nil, errors.New("Code: " + fmt.Sprint(ar.Code) + " Message:" + ar.Message)
	}

	return &ar, nil
}

func (c *Client) ModifyVs(v *Vs) (*Vs, error) {
	vsa := &VsApiPayLoad{
		Vs{
			Index:         v.Index,
			Address:       v.Address,
			Port:          v.Port,
			NickName:      v.NickName,
			SSLReverse:    v.SSLReverse,
			SSLReencrypt:  v.SSLReencrypt,
			InterceptMode: v.InterceptMode,
			Intercept:     v.Intercept,
			InterceptOpts: v.InterceptOpts,
			Enable:        v.Enable,
			ForceL4:       v.ForceL4,
			ForceL7:       v.ForceL7,
			Type:          v.Type,
			Protocol:      v.Protocol,
		},
		ApiPayLoad{
			ApiKey: c.ApiKey,
			CMD:    "modvs",
		},
	}

	b, err := json.Marshal(vsa)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.RestUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vs Vs
	err = json.Unmarshal(resp, &vs)
	if err != nil {
		return nil, err
	}

	return &vs, nil
}
