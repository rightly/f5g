package bigip

import (
	"encoding/json"
	"fmt"
)

type TopologyRecordList struct {
	Kind  string           `json:"kind,omitempty"`
	Items []TopologyRecord `json:"items,omitempty"`
}

type TopologyRecord struct {
	Kind       string `json:"kind,omitempty"`
	Name       string `json:"name,omitempty"`
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	Order      int    `json:"order,omitempty"`
	Score      int    `json:"score,omitempty"`
}

type topologyRecordId struct {
	RequestSourceType  string
	RequestSourceValue string
	DestinationType    string
	DestinationValue   string
}

func (g *GTM) TopologyRecordList() (*TopologyRecordList, error) {
	url := g.c.buildUrl(basePath, topologyResource)
	resp := &TopologyRecordList{
		Items: make([]TopologyRecord, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) TopologyRecord(id *topologyRecordId) (*TopologyRecord, error) {
	url := g.c.buildUrl(basePath, topologyResource, IdEncoding(id.TopologyId()))
	resp := &TopologyRecord{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewTopologyRecordId(requestSourceType string, requestSourceValue string, destinationType string, destinationValue string) *topologyRecordId {
	return &topologyRecordId{RequestSourceType: requestSourceType, RequestSourceValue: requestSourceValue, DestinationType: destinationType, DestinationValue: destinationValue}
}

func (t *topologyRecordId) TopologyId() string {
	return fmt.Sprintf("ldns: %s %s server: %s %s", t.RequestSourceType, t.RequestSourceValue, t.DestinationType, t.DestinationValue)
}

func NewTolopogyRecord() *TopologyRecord {
	return new(TopologyRecord)
}

// SetName topology resource name 설정
func (r *TopologyRecord) SetName(id *topologyRecordId) *TopologyRecord {
	r.Name = id.TopologyId()
	return r
}

func (r *TopologyRecord) SetWeight(weight int) *TopologyRecord {
	r.Score = weight
	return r
}

func (r *TopologyRecord) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set Monitor.Name")
	}
	return nil
}

func (g *GTM) CreateTopologyRecord(topologyConfig *TopologyRecord) (*TopologyRecord, error) {
	url := g.c.buildUrl(basePath, topologyResource)
	if err := topologyConfig.Verify(); err != nil {
		return nil, newError(400, "topologyRecord value verify fail: "+err.Error())
	}

	resp := &TopologyRecord{}
	body, err := json.Marshal(topologyConfig)
	if err != nil {
		return nil, newError(500, "CreateTopologyRecord.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(httpPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DeleteTopologyRecord(id *topologyRecordId) (*TopologyRecord, error) {
	url := g.c.buildUrl(basePath, topologyResource, IdEncoding(id.TopologyId()))

	resp := &TopologyRecord{}
	err := g.c.iControlRequest(httpDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
