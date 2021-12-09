package bigip

import (
	"encoding/json"
	"fmt"
)

type DatacenterList struct {
	Kind  string       `json:"kind,omitempty"`
	Items []Datacenter `json:"items,omitempty"`
}

type Datacenter struct {
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name"`
	Partition   string `json:"partition,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	Generation  int    `json:"generation,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Location    string `json:"location,omitempty"`
	Description string `json:"description,omitempty"`
}

func (g *GTM) DatacenterList() (*DatacenterList, error) {
	url := g.c.buildUrl(basePath, datacenterResource)
	resp := &DatacenterList{
		Items: make([]Datacenter, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) Datacenter(id string) (*Datacenter, error) {
	url := g.c.buildUrl(basePath, datacenterResource, CommonId(id))
	resp := &Datacenter{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DatacenterWithCustomId(id string) (*Datacenter, error) {
	url := g.c.buildUrl(basePath, datacenterResource, id)
	resp := &Datacenter{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewDatacenterConfig() *Datacenter {
	return new(Datacenter)
}

func (r *Datacenter) SetName(name string) *Datacenter {
	r.Name = name
	return r
}

func (r *Datacenter) SetLocation(location string) *Datacenter {
	r.Location = location
	return r
}

func (r *Datacenter) SetEnabled(state bool) *Datacenter {
	r.Enabled = state
	return r
}

func (r *Datacenter) SetPartition(partition string) *Datacenter {
	r.Partition = partition
	return r
}

func (r *Datacenter) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set Datacenter.Name")
	}
	if r.Partition == "" {
		r.SetPartition("Common")
	}
	return nil
}

func (g *GTM) CreateDatacenter(datacenterConfig *Datacenter) (*Datacenter, error) {
	url := g.c.buildUrl(basePath, datacenterResource)
	if err := datacenterConfig.Verify(); err != nil {
		return nil, newError(400, "datacenter value verify fail: "+err.Error())
	}

	resp := &Datacenter{}
	body, err := json.Marshal(datacenterConfig)
	if err != nil {
		return nil, newError(500, "CreateDatacenter.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(httpPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DeleteDatacenter(datacenterId string) (*Datacenter, error) {
	url := g.c.buildUrl(basePath, datacenterResource, CommonId(datacenterId))

	resp := &Datacenter{}
	err := g.c.iControlRequest(httpDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
