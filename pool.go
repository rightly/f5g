package f5g

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
	Kind                     string `json:"kind,omitempty"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	AlternateMode            string `json:"alternateMode,omitempty"`
	DynamicRatio             string `json:"dynamicRatio,omitempty"`
	Enabled                  bool   `json:"enabled,omitempty"`
	FallbackMode             string `json:"fallbackMode,omitempty"`
	FallbackIp               string `json:"fallbackIp,omitempty"`
	LoadBalancingMode        string `json:"loadBalancingMode,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	MaxAnswersReturned       int    `json:"maxAnswersReturned,omitempty"`
	QosHitRatio              int    `json:"qosHitRatio,omitempty"`
	QosHops                  int    `json:"qosHops,omitempty"`
	QosKilobytesSecond       int    `json:"qosKilobytesSecond,omitempty"`
	QosLcs                   int    `json:"qosLcs,omitempty"`
	QosPacketRate            int    `json:"qosPacketRate,omitempty"`
	QosRtt                   int    `json:"qosRtt,omitempty"`
	QosTopology              int    `json:"qosTopology,omitempty"`
	QosVsCapacity            int    `json:"qosVsCapacity,omitempty"`
	QosVsScore               int    `json:"qosVsScore,omitempty"`
	TTL                      int    `json:"ttl,omitempty"`
	VerifyMemberAvailability string `json:"verifyMemberAvailability,omitempty"`
	MembersReference         struct {
		Link            string `json:"link,omitempty"`
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
	} `json:"membersReference,omitempty"`
}

type PoolMemberList struct {
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selfLink,omitempty"`
	Items    []PoolMember `json:"items,omitempty"`
}

type PoolMember struct {
	Kind                      string `json:"kind,omitempty"`
	Name                      string `json:"name"`
	Partition                 string `json:"partition"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	SelfLink                  string `json:"selfLink,omitempty"`
	Enabled                   bool   `json:"enabled,omitempty"`
	StaticTarget              string `json:"staticTarget,omitempty"`
	LimitMaxBps               int    `json:"limitMaxBps,omitempty"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus,omitempty"`
	LimitMaxConnections       int    `json:"limitMaxConnections,omitempty"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus,omitempty"`
	LimitMaxPps               int    `json:"limitMaxPps,omitempty"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus,omitempty"`
	MemberOrder               int    `json:"memberOrder,omitempty"`
	Monitor                   string `json:"monitor,omitempty"`
	Ratio                     int    `json:"ratio,omitempty"`
}

/* GET Resource */

func (g *GTM) PoolList(ResourceType string) (*PoolList, error) {
	url := g.c.buildUrl(basePath, poolResource, ResourceType)
	resp := &PoolList{
		Items: make([]Pool, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) Pool(id, ResourceType string) (*Pool, error) {
	url := g.c.buildUrl(basePath, poolResource, ResourceType, CommonId(id))
	resp := &Pool{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) PoolWithCustomId(id, ResourceType string) (*Pool, error) {
	url := g.c.buildUrl(basePath, poolResource, ResourceType, id)
	resp := &Pool{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) PoolMemberList(id, ResourceType string) (*PoolMemberList, error) {
	url := g.c.buildUrl(basePath, poolResource, ResourceType, CommonId(id), memberResource)
	resp := &PoolMemberList{
		Items: make([]PoolMember, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) PoolMemberListWithCustomId(id, ResourceType string) (*PoolMemberList, error) {
	url := g.c.buildUrl(basePath, poolResource, ResourceType, id, memberResource)
	resp := &PoolMemberList{
		Items: make([]PoolMember, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
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
	url := g.c.buildUrl(basePath, poolResource, resourceType)
	if err := poolConfig.Verify(); err != nil {
		return nil, newError(400, "pool value verify fail: "+err.Error())
	}

	resp := &Pool{}
	body, err := json.Marshal(poolConfig)
	if err != nil {
		return nil, newError(500, "CreatePool.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(httpPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DeletePool(id, resourceType string) (*Pool, error) {
	url := g.c.buildUrl(basePath, poolResource, resourceType, CommonId(id))

	resp := &Pool{}
	err := g.c.iControlRequest(httpDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewPoolMemberConfig() *PoolMember {
	return new(PoolMember)
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

func (g *GTM) AddPoolMember(memberConfig *PoolMember, poolId, resourceType string) (*PoolMember, error) {
	url := g.c.buildUrl(basePath, poolResource, resourceType, poolId, memberResource)
	if err := memberConfig.Verify(); err != nil {
		return nil, newError(400, "pool value verify fail: "+err.Error())
	}

	resp := &PoolMember{}
	body, err := json.Marshal(memberConfig)
	if err != nil {
		return nil, newError(500, "CreatePool.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(httpPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) RemovePoolMember(poolId, poolMemberId, resourceType string) (*PoolMember, error) {
	if !strings.Contains(poolMemberId, ":") {
		poolMemberId = poolMemberId + ":" + poolMemberId
	}
	url := g.c.buildUrl(basePath, poolResource, resourceType, CommonId(poolId), memberResource, CommonId(poolMemberId))

	resp := &PoolMember{}
	err := g.c.iControlRequest(httpDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
