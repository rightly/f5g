package main

import (
	"encoding/json"
	"fmt"
)

type MonitorList struct {
	Kind     string    `json:"kind"`
	SelfLink string    `json:"selfLink"`
	Items    []Monitor `json:"items"`
}

type Monitor struct {
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	DefaultsFrom       string `json:"defaultsFrom,omitempty"`
	Destination        string `json:"destination,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	Reverse            string `json:"reverse,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Transparent        string `json:"transparent,omitempty"`
	Send               string `json:"send,omitempty"`
	Receive            string `json:"recv,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Cipherlist         string `json:"cipherlist,omitempty"`
	Compatibility      string `json:"compatibility,omitempty"`
}

/* GET Resource */

// MonitorList 모니터 리스트 응답
// monitorType: 모니터 프로토콜 타입 (bigip.MonitorHTTP ... )
func (g *GTM) MonitorList(monitorType string) (*MonitorList, error) {
	url := g.c.buildUrl(basePath, monitorResource, monitorType)
	resp := &MonitorList{
		Items: make([]Monitor, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Monitor 모니터 정보 응답
// id: 모니터 id
// monitorType: 모니터 프로토콜 타입 (bigip.MonitorHTTP ... )
func (g *GTM) Monitor(id, monitorType string) (*Monitor, error) {
	url := g.c.buildUrl(basePath, monitorResource, monitorType, CommonId(id))
	resp := &Monitor{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// MonitorWithCustomId Common partition 이 아닌 id 로 모니터 정보 확인
// id: 모니터 id (with partition e.g: ~Custom~monitorId)
// monitorType: 모니터 프로토콜 타입 (bigip.MonitorHTTP ... )
func (g *GTM) MonitorWithCustomId(id, monitorType string) (*Monitor, error) {
	url := g.c.buildUrl(basePath, monitorResource, monitorType, id)
	resp := &Monitor{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewMonitorConfig() *Monitor {
	return new(Monitor)
}

func (r *Monitor) SetName(name string) *Monitor {
	r.Name = name
	return r
}

func (r *Monitor) SetPartition(partition string) *Monitor {
	r.Partition = partition
	return r
}

func (r *Monitor) SetInterval(interval int) *Monitor {
	r.Interval = interval
	return r
}

func (r *Monitor) SetTimeout(timeout int) *Monitor {
	r.Timeout = timeout
	return r
}

func (r *Monitor) SetProbeTimeout(timeout int) *Monitor {
	r.ProbeTimeout = timeout
	return r
}

func (r *Monitor) SetSendString(send string) *Monitor {
	r.Send = send
	return r
}

func (r *Monitor) SetReceiveString(recv string) *Monitor {
	r.Receive = recv
	return r
}

func (r *Monitor) SetChiperList(cipher string) *Monitor {
	r.Cipherlist = cipher
	return r
}

func (r *Monitor) SetUsername(username string) *Monitor {
	r.Username = username
	return r
}

func (r *Monitor) SetPassword(passwd string) *Monitor {
	r.Password = passwd
	return r
}

func (r *Monitor) SetAliasDestination(ip, port string) *Monitor {
	r.Destination = ip + ":" + port
	return r
}

func (r *Monitor) SetIgnoreDownResponse(ignore bool) *Monitor {
	r.IgnoreDownResponse = "disabled"
	if ignore {
		r.IgnoreDownResponse = "enabled"
	}

	return r
}

func (r *Monitor) SetReverse(reverse bool) *Monitor {
	r.Reverse = "disabled"
	if reverse {
		r.Reverse = "enabled"
	}

	return r
}

func (r *Monitor) Verify(monitorType string) error {
	if r.Name == "" {
		return fmt.Errorf("please set Monitor.Name")
	}
	if r.Partition == "" {
		r.SetPartition("Common")
	}
	if monitorType == MonitorHTTPS {
		if r.Compatibility == "" {
			return fmt.Errorf("please set Monitor.Compatibility")
		}
	}

	return nil
}

func (g *GTM) CreateMonitor(monitorConfig *Monitor, monitorType string) (*Monitor, error) {
	url := g.c.buildUrl(basePath, monitorResource, monitorType)
	if err := monitorConfig.Verify(monitorType); err != nil {
		return nil, newError(400, "monitor value verify fail: "+err.Error())
	}

	resp := &Monitor{}
	body, err := json.Marshal(monitorConfig)
	if err != nil {
		return nil, newError(500, "CreateMonitor.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(httpPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DeleteMonitor(monitorId, monitorType string) (*Monitor, error) {
	url := g.c.buildUrl(basePath, monitorResource, monitorType, CommonId(monitorId))

	resp := &Monitor{}
	err := g.c.iControlRequest(httpDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
