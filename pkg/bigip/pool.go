package bigip

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PoolList struct {
	Kind  string `json:"kind"`
	Items []Pool `json:"items"`
}

type Pool struct {
	Kind                     string `json:"kind"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	SelfLink                 string `json:"selfLink"`
	AlternateMode            string `json:"alternateMode"`
	DynamicRatio             string `json:"dynamicRatio"`
	Enabled                  bool   `json:"enabled"`
	FallbackMode             string `json:"fallbackMode"`
	FallbackIp               string `json:"fallbackIp"`
	LoadBalancingMode        string `json:"loadBalancingMode"`
	ManualResume             string `json:"manualResume"`
	MaxAnswersReturned       int    `json:"maxAnswersReturned"`
	QosHitRatio              int    `json:"qosHitRatio"`
	QosHops                  int    `json:"qosHops"`
	QosKilobytesSecond       int    `json:"qosKilobytesSecond"`
	QosLcs                   int    `json:"qosLcs"`
	QosPacketRate            int    `json:"qosPacketRate"`
	QosRtt                   int    `json:"qosRtt"`
	QosTopology              int    `json:"qosTopology"`
	QosVsCapacity            int    `json:"qosVsCapacity"`
	QosVsScore               int    `json:"qosVsScore"`
	TTL                      int    `json:"ttl"`
	VerifyMemberAvailability string `json:"verifyMemberAvailability"`
	MembersReference         struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"membersReference"`
}

type PoolMemberList struct {
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selfLink"`
	Items    []PoolMember `json:"items"`
}

type PoolMember struct {
	Kind                      string `json:"kind"`
	Name                      string `json:"name"`
	Partition                 string `json:"partition"`
	FullPath                  string `json:"fullPath"`
	Generation                int    `json:"generation"`
	SelfLink                  string `json:"selfLink"`
	Enabled                   bool   `json:"enabled"`
	StaticTarget              string `json:"staticTarget"`
	LimitMaxBps               int    `json:"limitMaxBps"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus"`
	LimitMaxConnections       int    `json:"limitMaxConnections"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus"`
	LimitMaxPps               int    `json:"limitMaxPps"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus"`
	MemberOrder               int    `json:"memberOrder"`
	Monitor                   string `json:"monitor"`
	Ratio                     int    `json:"ratio"`
}

/* GET Resource */

func (g *GTM) PoolList(ResourceType string) (*PoolList, error) {
	url := g.c.buildUrl(basePath, PoolResource, ResourceType)
	resp := &PoolList{
		Items: make([]Pool, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) Pool(id, ResourceType string) (*Pool, error) {
	url := g.c.buildUrl(basePath, PoolResource, ResourceType, CommonId(id))
	resp := &Pool{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) PoolWithCustomId(id, ResourceType string) (*Pool, error) {
	url := g.c.buildUrl(basePath, PoolResource, ResourceType, id)
	resp := &Pool{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) PoolMemberList(id, ResourceType string) (*PoolMemberList, error) {
	url := g.c.buildUrl(basePath, PoolResource, ResourceType, CommonId(id), MemberResource)
	resp := &PoolMemberList{
		Items: make([]PoolMember, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) PoolMemberListWithCustomId(id, ResourceType string) (*PoolMemberList, error) {
	url := g.c.buildUrl(basePath, PoolResource, ResourceType, id, MemberResource)
	resp := &PoolMemberList{
		Items: make([]PoolMember, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/* Modify Resource */

func NewPoolConfig() *Pool {
	return new(Pool)
}

func (r *Pool) SetName(name string) *Pool {
	r.Name = name
	return r
}

func (r *Pool) SetTTL(ttl int) *Pool {
	r.TTL = ttl
	return r
}

func (r *Pool) SetVerifyMemberAvailability(b bool) *Pool {
	r.VerifyMemberAvailability = "enabled"
	if !b {
		r.VerifyMemberAvailability = "disabled"
	}

	return r
}

func (r *Pool) SetMonitor(monitor ...string) *Pool {
	i := 0
	fs := " and "
	for _, v := range monitor {
		if strings.Contains(v, "/") {
			v = fmt.Sprintf("/Common/%s", v)
		}
		if i < len(monitor)-1 {
			r.Name += v + fs
		}
	}

	return r
}

func (r *Pool) SetMaxAnswersReturned(answer int) *Pool {
	r.MaxAnswersReturned = answer
	return r
}

func (r *Pool) SetLoadBalancingMode(lbMode string) *Pool {
	r.LoadBalancingMode = lbMode
	return r
}

func (r *Pool) SetAlternateMode(alMode string) *Pool {
	r.AlternateMode = alMode
	return r
}

func (r *Pool) SetFallbackMode(fbMode string) *Pool {
	r.FallbackMode = fbMode
	return r
}

func (r *Pool) SetFallbackIp(ip string) *Pool {
	r.FallbackIp = ip
	return r
}

func (r *Pool) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set Pool.Name")
	}

	return nil
}

func (g *GTM) CreatePool(poolConfig *Pool, resourceType string) (*Pool, error) {
	url := g.c.buildUrl(basePath, PoolResource, resourceType)
	if err := poolConfig.Verify(); err != nil {
		return nil, newError(400, "pool value verify fail: "+err.Error())
	}

	resp := &Pool{}
	body, err := json.Marshal(poolConfig)
	if err != nil {
		return nil, newError(500, "CreatePool.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(HTTPPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DeletePool(id, resourceType string) (*Pool, error) {
	url := g.c.buildUrl(basePath, PoolResource, resourceType, CommonId(id))

	resp := &Pool{}
	err := g.c.iControlRequest(HTTPDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *PoolMember) SetName(name string) *PoolMember {
	r.Name = name
	if !strings.Contains(name, ":") {
		r.Name = name + ":" + name
	}
	return r
}

func (r *PoolMember) SetPartition(partition string) *PoolMember {
	r.Partition = partition
	return r
}

func (r *PoolMember) SetStaticTarget(static bool) *PoolMember {
	r.StaticTarget = "no"
	if static {
		r.StaticTarget = "yes"
	}

	return r
}

func (r *PoolMember) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set PoolMember.Name")
	}
	if r.Partition == "" {
		r.Partition = "Common"
	}

	return nil
}

func (g *GTM) AddPoolMember(memberConfig *PoolMember, id, resourceType string) (*PoolMember, error) {
	url := g.c.buildUrl(basePath, PoolResource, resourceType, id, MemberResource)
	if err := memberConfig.Verify(); err != nil {
		return nil, newError(400, "pool value verify fail: "+err.Error())
	}

	resp := &PoolMember{}
	body, err := json.Marshal(memberConfig)
	if err != nil {
		return nil, newError(500, "CreatePool.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(HTTPPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) RemovePool(poolId, poolMemberId, resourceType string) (*PoolMember, error) {
	if !strings.Contains(poolMemberId, ":") {
		poolMemberId = poolMemberId + ":" + poolMemberId
	}
	url := g.c.buildUrl(basePath, PoolResource, resourceType, CommonId(poolId), MemberResource, CommonId(poolMemberId))

	resp := &PoolMember{}
	err := g.c.iControlRequest(HTTPDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
