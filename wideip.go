package f5g

import (
	"encoding/json"
	"fmt"
	"strings"
)

// WideIpList wideip 목록 struct
type WideIpList struct {
	Kind  string   `json:"kind"`
	Items []WideIp `json:"items"`
}

// WideIp wideip struct
type WideIp struct {
	Kind                 string       `json:"kind,omitempty"`
	Name                 string       `json:"name"`
	Partition            string       `json:"partition"`
	FullPath             string       `json:"fullPath,omitempty"`
	Generation           int          `json:"generation,omitempty"`
	Enabled              bool         `json:"enabled,omitempty"`
	FailureRcode         string       `json:"failureRcode,omitempty"`
	FailureRcodeResponse string       `json:"failureRcodeResponse,omitempty"`
	FailureRcodeTTL      int          `json:"failureRcodeTtl,omitempty"`
	LastResortPool       string       `json:"lastResortPool,omitempty"`
	MinimalResponse      string       `json:"minimalResponse,omitempty"`
	PersistCidrIpv4      int          `json:"persistCidrIpv4,omitempty"`
	PersistCidrIpv6      int          `json:"persistCidrIpv6,omitempty"`
	Persistence          string       `json:"persistenc,omitemptye"`
	PoolLbMode           string       `json:"poolLbMode,omitempty"`
	TTLPersistence       int          `json:"ttlPersistence,omitempty"`
	Pools                []WideIpPool `json:"pools,omitempty"`
	PoolsCname           []WideIpPool `json:"poolsCname,omitempty"`
}

// WideIpPool wideip 의 pool 정보 struct
type WideIpPool struct {
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	Order         int    `json:"order,omitempty"`
	Ratio         int    `json:"ratio,omitempty"`
	NameReference struct {
		Link string `json:"link,omitempty"`
	} `json:"nameReference,omitempty"`
}

// WideIpList wide ip 목록 응답
func (g *GTM) WideIpList(ResourceType string) (*WideIpList, error) {
	url := g.c.buildUrl(basePath, wideIpResource, ResourceType)
	resp := &WideIpList{
		Items: make([]WideIp, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// WideIp 특정 id의 wide ip 응답
func (g *GTM) WideIp(id, ResourceType string) (*WideIp, error) {
	url := g.c.buildUrl(basePath, wideIpResource, ResourceType, CommonId(id))
	resp := &WideIp{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// WideIpWithCustomId Common 외 다른  partition 을 사용해 특정 ip의 wide ip 응답
func (g *GTM) WideIpWithCustomId(id, ResourceType string) (*WideIp, error) {
	url := g.c.buildUrl(basePath, wideIpResource, ResourceType, id)
	resp := &WideIp{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// NewWideIpConfig wide ip struct 생성
func NewWideIpConfig() *WideIp {
	return &WideIp{}
}

func (r *WideIp) SetName(name string) *WideIp {
	r.Name = name
	return r
}

func (r *WideIp) SetPoolLBMode(lbMode string) *WideIp {
	r.PoolLbMode = lbMode
	return r
}

func (r *WideIp) SetLastResortPool(poolId string) *WideIp {
	r.LastResortPool = poolId
	if !strings.Contains(r.LastResortPool, "/") {
		r.LastResortPool = fmt.Sprintf("/Common/%s", r.LastResortPool)
	}
	return r
}

func (r *WideIp) SetPartition(partition string) *WideIp {
	r.Partition = partition
	return r
}

func (r *WideIp) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set WideIp.Name")
	}
	if r.Partition == "" {
		r.SetPartition("Common")
	}
	return nil
}

// CreateWideIp wide ip 생성
func (g *GTM) CreateWideIp(wideipConfig *WideIp, resourceType string) (*WideIp, error) {
	url := g.c.buildUrl(basePath, wideIpResource, resourceType)
	if err := wideipConfig.Verify(); err != nil {
		return nil, newError(400, "wideIp value verify fail: "+err.Error())
	}

	resp := &WideIp{}
	body, err := json.Marshal(wideipConfig)
	if err != nil {
		return nil, newError(500, "CreateWideIp.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(httpPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteWideIp wide ip 삭제
func (g *GTM) DeleteWideIp(wideipId, resourceType string) (*WideIp, error) {
	url := g.c.buildUrl(basePath, wideIpResource, resourceType, CommonId(wideipId))

	resp := &WideIp{}
	err := g.c.iControlRequest(httpDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// NewWideIpPoolConfig wide ip pool struct 생성
func NewWideIpPoolConfig() *WideIpPool {
	return new(WideIpPool)
}

func (r *WideIpPool) SetName(name string) *WideIpPool {
	r.Name = name
	return r
}

func (r *WideIpPool) SetPartition(partition string) *WideIpPool {
	r.Partition = partition
	return r
}

func (r *WideIpPool) SetOrder(order int) *WideIpPool {
	r.Order = order
	return r
}

func (r *WideIpPool) SetRatio(ratio int) *WideIpPool {
	r.Ratio = ratio
	return r
}

func (r *WideIpPool) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set WideIpPool.Name")
	}
	if r.Partition == "" {
		r.SetPartition("Common")
	}
	return nil
}

// AddPools wide ip struct 에 pool 정보 추가
func (r *WideIp) AddPools(resourceType string, pools ...*WideIpPool) error {
	for _, pool := range pools {
		if err := pool.Verify(); err != nil {
			return newError(400, "wideIpPool value verify fail: "+err.Error())
		}

		switch resourceType {
		case CNAMEType:
			r.PoolsCname = append(r.PoolsCname, *pool)
		default:
			r.Pools = append(r.Pools, *pool)
		}
	}

	return nil
}

// UpdateWideIp wide ip 수정
func (g *GTM) UpdateWideIp(wideipConfig *WideIp, wideIpId, resourceType string) (*WideIp, error) {
	url := g.c.buildUrl(basePath, wideIpResource, resourceType, CommonId(wideIpId))
	if err := wideipConfig.Verify(); err != nil {
		return nil, newError(400, "wideIp value verify fail: "+err.Error())
	}

	resp := &WideIp{}
	body, err := json.Marshal(wideipConfig)
	if err != nil {
		return nil, newError(500, "CreateWideIp.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(httpPut, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
