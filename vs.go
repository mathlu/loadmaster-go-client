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
	Status                  string   `json:"Status,omitempty"`
	Index                   int      `json:"Index,omitempty"`
	Address                 string   `json:"VSAddress"`
	Port                    string   `json:"VSPort"`
	Layer                   int      `json:"Layer,omitempty"`
	NickName                string   `json:"NickName,omitempty"`
	Enable                  bool     `json:"Enable,omitempty"`
	SSLReverse              bool     `json:"SSLReverse,omitempty"`
	SSLReencrypt            bool     `json:"SSLReencrypt,omitempty"`
	InterceptMode           int      `json:"InterceptMode,omitempty"`
	Intercept               bool     `json:"Intercept,omitempty"`
	InterceptOpts           []string `json:"InterceptOpts,omitempty"`
	AlertThreshold          int      `json:"AlertThreshold,omitempty"`
	OwaspOpts               []string `json:"OwaspOpts,omitempty"`
	BlockingParanoia        int      `json:"BlockingParanoia,omitempty"`
	IPReputationBlocking    bool     `json:"IPReputationBlocking,omitempty"`
	ExecutingParanoia       int      `json:"ExecutingParanoia,omitempty"`
	AnomalyScoringThreshold int      `json:"AnomalyScoringThreshold,omitempty"`
	PCRELimit               int      `json:"PCRELimit,omitempty"`
	JSONDLimit              int      `json:"JSONDLimit,omitempty"`
	BodyLimit               int      `json:"BodyLimit,omitempty"`
	Transactionlimit        int      `json:"Transactionlimit,omitempty"`
	Transparent             bool     `json:"Transparent,omitempty"`
	SubnetOriginating       bool     `json:"SubnetOriginating,omitempty"`
	ServerInit              int      `json:"ServerInit,omitempty"`
	StartTLSMode            int      `json:"StartTLSMode,omitempty"`
	Idletime                int      `json:"Idletime,omitempty"`
	Cache                   bool     `json:"Cache,omitempty"`
	Compress                bool     `json:"Compress,omitempty"`
	Verify                  int      `json:"Verify,omitempty"`
	UseforSnat              bool     `json:"UseforSnat,omitempty"`
	ForceL4                 bool     `json:"ForceL4,omitempty"`
	ForceL7                 bool     `json:"ForceL7,omitempty"`
	MultiConnect            bool     `json:"MultiConnect,omitempty"`
	ClientCert              int      `json:"ClientCert,omitempty"`
	SecurityHeaderOptions   int      `json:"SecurityHeaderOptions,omitempty"`
	SameSite                int      `json:"SameSite,omitempty"`
	VerifyBearer            bool     `json:"VerifyBearer,omitempty"`
	ErrorCode               string   `json:"ErrorCode,omitempty"`
	CheckUse11              bool     `json:"CheckUse1.1,omitempty"`
	MatchLen                int      `json:"MatchLen,omitempty"`
	CheckUseGet             int      `json:"CheckUseGet,omitempty"`
	SSLRewrite              string   `json:"SSLRewrite,omitempty"`
	Type                    string   `json:"VStype,omitempty"`
	FollowVSID              int      `json:"FollowVSID,omitempty"`
	Protocol                string   `json:"Protocol"`
	Schedule                string   `json:"Schedule,omitempty"`
	CheckType               string   `json:"CheckType,omitempty"`
	PersistTimeout          string   `json:"PersistTimeout,omitempty"`
	CheckPort               string   `json:"CheckPort,omitempty"`
	HTTPReschedule          bool     `json:"HTTPReschedule,omitempty"`
	NRules                  int      `json:"NRules,omitempty"`
	NRequestRules           int      `json:"NRequestRules,omitempty"`
	NResponseRules          int      `json:"NResponseRules,omitempty"`
	NMatchBodyRules         int      `json:"NMatchBodyRules,omitempty"`
	NPreProcessRules        int      `json:"NPreProcessRules,omitempty"`
	EspEnabled              bool     `json:"EspEnabled,omitempty"`
	InputAuthMode           int      `json:"InputAuthMode,omitempty"`
	OutputAuthMode          int      `json:"OutputAuthMode,omitempty"`
	MasterVS                int      `json:"MasterVS,omitempty"`
	MasterVSID              int      `json:"MasterVSID,omitempty"`
	IsTransparent           int      `json:"IsTransparent,omitempty"`
	AddVia                  int      `json:"AddVia,omitempty"`
	QoS                     int      `json:"QoS,omitempty"`
	TLSType                 string   `json:"TlsType,omitempty"`
	NeedHostName            bool     `json:"NeedHostName,omitempty"`
	OCSPVerify              bool     `json:"OCSPVerify,omitempty"`
	AllowHTTP2              bool     `json:"AllowHTTP2,omitempty"`
	PassCipher              bool     `json:"PassCipher,omitempty"`
	PassSni                 bool     `json:"PassSni,omitempty"`
	ChkInterval             int      `json:"ChkInterval,omitempty"`
	ChkTimeout              int      `json:"ChkTimeout,omitempty"`
	ChkRetryCount           int      `json:"ChkRetryCount,omitempty"`
	Bandwidth               int      `json:"Bandwidth,omitempty"`
	ConnsPerSecLimit        int      `json:"ConnsPerSecLimit,omitempty"`
	RequestsPerSecLimit     int      `json:"RequestsPerSecLimit,omitempty"`
	MaxConnsLimit           int      `json:"MaxConnsLimit,omitempty"`
	RefreshPersist          bool     `json:"RefreshPersist,omitempty"`
	EnhancedHealthChecks    bool     `json:"EnhancedHealthChecks,omitempty"`
	RsMinimum               int      `json:"RsMinimum,omitempty"`
	NumberOfRSs             int      `json:"NumberOfRSs,omitempty"`
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
			Port                    string `json:"vsport"`
			NickName                string `json:"NickName,omitempty"`
			Enable                  bool   `json:"Enable,omitempty"`
			SSLReverse              bool   `json:"SSLReverse,omitempty"`
			SSLReencrypt            bool   `json:"SSLReencrypt,omitempty"`
			InterceptMode           int    `json:"InterceptMode,omitempty"`
			Intercept               bool   `json:"Intercept,omitempty"`
			InterceptOpts           string `json:"InterceptOpts,omitempty"`
			AlertThreshold          int    `json:"AlertThreshold,omitempty"`
			OwaspOpts               string `json:"OwaspOpts,omitempty"`
			BlockingParanoia        int    `json:"BlockingParanoia,omitempty"`
			IPReputationBlocking    bool   `json:"IPReputationBlocking,omitempty"`
			ExecutingParanoia       int    `json:"ExecutingParanoia,omitempty"`
			AnomalyScoringThreshold int    `json:"AnomalyScoringThreshold,omitempty"`
			PCRELimit               int    `json:"PCRELimit,omitempty"`
			JSONDLimit              int    `json:"JSONDLimit,omitempty"`
			BodyLimit               int    `json:"BodyLimit,omitempty"`
			Transactionlimit        int    `json:"Transactionlimit,omitempty"`
			Transparent             bool   `json:"Transparent,omitempty"`
			SubnetOriginating       bool   `json:"SubnetOriginating,omitempty"`
			ServerInit              int    `json:"ServerInit,omitempty"`
			StartTLSMode            int    `json:"StartTLSMode,omitempty"`
			Idletime                int    `json:"Idletime,omitempty"`
			Cache                   bool   `json:"Cache,omitempty"`
			Compress                bool   `json:"Compress,omitempty"`
			Verify                  int    `json:"Verify,omitempty"`
			UseforSnat              bool   `json:"UseforSnat,omitempty"`
			ForceL4                 bool   `json:"ForceL4,omitempty"`
			ForceL7                 bool   `json:"ForceL7,omitempty"`
			MultiConnect            bool   `json:"MultiConnect,omitempty"`
			ClientCert              int    `json:"ClientCert,omitempty"`
			SecurityHeaderOptions   int    `json:"SecurityHeaderOptions,omitempty"`
			SameSite                int    `json:"SameSite,omitempty"`
			VerifyBearer            bool   `json:"VerifyBearer,omitempty"`
			ErrorCode               string `json:"ErrorCode,omitempty"`
			CheckUse11              bool   `json:"CheckUse1.1,omitempty"`
			MatchLen                int    `json:"MatchLen,omitempty"`
			CheckUseGet             int    `json:"CheckUseGet,omitempty"`
			SSLRewrite              string `json:"SSLRewrite,omitempty"`
			Type                    string `json:"VStype,omitempty"`
			FollowVSID              int    `json:"FollowVSID,omitempty"`
			Protocol                string `json:"prot,omitempty"`
			Schedule                string `json:"Schedule,omitempty"`
			CheckType               string `json:"CheckType,omitempty"`
			PersistTimeout          string `json:"PersistTimeout,omitempty"`
			CheckPort               string `json:"CheckPort,omitempty"`
			HTTPReschedule          bool   `json:"HTTPReschedule,omitempty"`
			NRules                  int    `json:"NRules,omitempty"`
			NRequestRules           int    `json:"NRequestRules,omitempty"`
			NResponseRules          int    `json:"NResponseRules,omitempty"`
			NMatchBodyRules         int    `json:"NMatchBodyRules,omitempty"`
			NPreProcessRules        int    `json:"NPreProcessRules,omitempty"`
			EspEnabled              bool   `json:"EspEnabled,omitempty"`
			InputAuthMode           int    `json:"InputAuthMode,omitempty"`
			OutputAuthMode          int    `json:"OutputAuthMode,omitempty"`
			MasterVS                int    `json:"MasterVS,omitempty"`
			MasterVSID              int    `json:"MasterVSID,omitempty"`
			IsTransparent           int    `json:"IsTransparent,omitempty"`
			AddVia                  int    `json:"AddVia,omitempty"`
			QoS                     int    `json:"QoS,omitempty"`
			TLSType                 string `json:"TlsType,omitempty"`
			NeedHostName            bool   `json:"NeedHostName,omitempty"`
			OCSPVerify              bool   `json:"OCSPVerify,omitempty"`
			AllowHTTP2              bool   `json:"AllowHTTP2,omitempty"`
			PassCipher              bool   `json:"PassCipher,omitempty"`
			PassSni                 bool   `json:"PassSni,omitempty"`
			ChkInterval             int    `json:"ChkInterval,omitempty"`
			ChkTimeout              int    `json:"ChkTimeout,omitempty"`
			ChkRetryCount           int    `json:"ChkRetryCount,omitempty"`
			Bandwidth               int    `json:"Bandwidth,omitempty"`
			ConnsPerSecLimit        int    `json:"ConnsPerSecLimit,omitempty"`
			RequestsPerSecLimit     int    `json:"RequestsPerSecLimit,omitempty"`
			MaxConnsLimit           int    `json:"MaxConnsLimit,omitempty"`
			RefreshPersist          bool   `json:"RefreshPersist,omitempty"`
			EnhancedHealthChecks    bool   `json:"EnhancedHealthChecks,omitempty"`
			RsMinimum               int    `json:"RsMinimum,omitempty"`
			NumberOfRSs             int    `json:"NumberOfRSs,omitempty"`
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
			NickName                string `json:"NickName,omitempty"`
			Enable                  bool   `json:"Enable,omitempty"`
			SSLReverse              bool   `json:"SSLReverse,omitempty"`
			SSLReencrypt            bool   `json:"SSLReencrypt,omitempty"`
			InterceptMode           int    `json:"InterceptMode,omitempty"`
			Intercept               bool   `json:"Intercept,omitempty"`
			InterceptOpts           string `json:"InterceptOpts,omitempty"`
			AlertThreshold          int    `json:"AlertThreshold,omitempty"`
			OwaspOpts               string `json:"OwaspOpts,omitempty"`
			BlockingParanoia        int    `json:"BlockingParanoia,omitempty"`
			IPReputationBlocking    bool   `json:"IPReputationBlocking,omitempty"`
			ExecutingParanoia       int    `json:"ExecutingParanoia,omitempty"`
			AnomalyScoringThreshold int    `json:"AnomalyScoringThreshold,omitempty"`
			PCRELimit               int    `json:"PCRELimit,omitempty"`
			JSONDLimit              int    `json:"JSONDLimit,omitempty"`
			BodyLimit               int    `json:"BodyLimit,omitempty"`
			Transactionlimit        int    `json:"Transactionlimit,omitempty"`
			Transparent             bool   `json:"Transparent,omitempty"`
			SubnetOriginating       bool   `json:"SubnetOriginating,omitempty"`
			ServerInit              int    `json:"ServerInit,omitempty"`
			StartTLSMode            int    `json:"StartTLSMode,omitempty"`
			Idletime                int    `json:"Idletime,omitempty"`
			Cache                   bool   `json:"Cache,omitempty"`
			Compress                bool   `json:"Compress,omitempty"`
			Verify                  int    `json:"Verify,omitempty"`
			UseforSnat              bool   `json:"UseforSnat,omitempty"`
			ForceL4                 bool   `json:"ForceL4,omitempty"`
			ForceL7                 bool   `json:"ForceL7,omitempty"`
			MultiConnect            bool   `json:"MultiConnect,omitempty"`
			ClientCert              int    `json:"ClientCert,omitempty"`
			SecurityHeaderOptions   int    `json:"SecurityHeaderOptions,omitempty"`
			SameSite                int    `json:"SameSite,omitempty"`
			VerifyBearer            bool   `json:"VerifyBearer,omitempty"`
			ErrorCode               string `json:"ErrorCode,omitempty"`
			CheckUse11              bool   `json:"CheckUse1.1,omitempty"`
			MatchLen                int    `json:"MatchLen,omitempty"`
			CheckUseGet             int    `json:"CheckUseGet,omitempty"`
			SSLRewrite              string `json:"SSLRewrite,omitempty"`
			Type                    string `json:"VStype,omitempty"`
			FollowVSID              int    `json:"FollowVSID,omitempty"`
			Protocol                string `json:"prot,omitempty"`
			Schedule                string `json:"Schedule,omitempty"`
			CheckType               string `json:"CheckType,omitempty"`
			PersistTimeout          string `json:"PersistTimeout,omitempty"`
			CheckPort               string `json:"CheckPort,omitempty"`
			HTTPReschedule          bool   `json:"HTTPReschedule,omitempty"`
			NRules                  int    `json:"NRules,omitempty"`
			NRequestRules           int    `json:"NRequestRules,omitempty"`
			NResponseRules          int    `json:"NResponseRules,omitempty"`
			NMatchBodyRules         int    `json:"NMatchBodyRules,omitempty"`
			NPreProcessRules        int    `json:"NPreProcessRules,omitempty"`
			EspEnabled              bool   `json:"EspEnabled,omitempty"`
			InputAuthMode           int    `json:"InputAuthMode,omitempty"`
			OutputAuthMode          int    `json:"OutputAuthMode,omitempty"`
			MasterVS                int    `json:"MasterVS,omitempty"`
			MasterVSID              int    `json:"MasterVSID,omitempty"`
			IsTransparent           int    `json:"IsTransparent,omitempty"`
			AddVia                  int    `json:"AddVia,omitempty"`
			QoS                     int    `json:"QoS,omitempty"`
			TLSType                 string `json:"TlsType,omitempty"`
			NeedHostName            bool   `json:"NeedHostName,omitempty"`
			OCSPVerify              bool   `json:"OCSPVerify,omitempty"`
			AllowHTTP2              bool   `json:"AllowHTTP2,omitempty"`
			PassCipher              bool   `json:"PassCipher,omitempty"`
			PassSni                 bool   `json:"PassSni,omitempty"`
			ChkInterval             int    `json:"ChkInterval,omitempty"`
			ChkTimeout              int    `json:"ChkTimeout,omitempty"`
			ChkRetryCount           int    `json:"ChkRetryCount,omitempty"`
			Bandwidth               int    `json:"Bandwidth,omitempty"`
			ConnsPerSecLimit        int    `json:"ConnsPerSecLimit,omitempty"`
			RequestsPerSecLimit     int    `json:"RequestsPerSecLimit,omitempty"`
			MaxConnsLimit           int    `json:"MaxConnsLimit,omitempty"`
			RefreshPersist          bool   `json:"RefreshPersist,omitempty"`
			EnhancedHealthChecks    bool   `json:"EnhancedHealthChecks,omitempty"`
			RsMinimum               int    `json:"RsMinimum,omitempty"`
			NumberOfRSs             int    `json:"NumberOfRSs,omitempty"`
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
