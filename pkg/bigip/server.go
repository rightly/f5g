package bigip

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ServerList struct {
	Kind  string   `json:"kind,omitempty"`
	Items []Server `json:"items,omitempty"`
}

type Server struct {
	Kind       string    `json:"kind,omitempty"`
	Name       string    `json:"name,omitempty"`
	Partition  string    `json:"partition,omitempty"`
	FullPath   string    `json:"fullPath,omitempty"`
	Generation int       `json:"generation,omitempty"`
	Datacenter string    `json:"datacenter,omitempty"`
	Enabled    bool      `json:"enabled,omitempty"`
	Monitor    string    `json:"monitor,omitempty"`
	Product    string    `json:"product,omitempty"`
	Addresses  []Address `json:"addresses,omitempty"`
}

type Address struct {
	Name        string `json:"name,omitempty"`
	DeviceName  string `json:"deviceName,omitempty"`
	Translation string `json:"translation,omitempty"`
}

type VirtualServerList struct {
	Kind  string          `json:"kind,omitempty"`
	Items []VirtualServer `json:"items,omitempty"`
}

type VirtualServer struct {
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	Destination        string `json:"destination,omitempty"`
	Enabled            bool   `json:"enabled,omitempty"`
	TranslationAddress string `json:"translationAddress,omitempty"`
	TranslationPort    int    `json:"translationPort,omitempty"`
}

/* GET Resource */

// ServerList 서버 목록 응답
func (g *GTM) ServerList() (*ServerList, error) {
	url := g.c.buildUrl(basePath, ServerResource)
	resp := &ServerList{
		Items: make([]Server, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Server resource id로 서버 응답
func (g *GTM) Server(id string) (*Server, error) {
	url := g.c.buildUrl(basePath, ServerResource, CommonId(id))
	resp := &Server{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ServerWithCustomId Common 이 아닌 별도 Partition 으로 resource id 서버 검색
func (g *GTM) ServerWithCustomId(id string) (*Server, error) {
	url := g.c.buildUrl(basePath, ServerResource, id)
	resp := &Server{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// VirtualServerList Virtual server 목록 응답
func (g *GTM) VirtualServerList(id string) (*VirtualServer, error) {
	url := g.c.buildUrl(basePath, ServerResource, CommonId(id), VirtualServerResource)
	resp := &VirtualServer{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// VirtualServer Virtual server 응답
func (g *GTM) VirtualServer(serverId, virtualServerId string) (*VirtualServer, error) {
	url := g.c.buildUrl(basePath, ServerResource, CommonId(serverId), VirtualServerResource, virtualServerId)
	resp := &VirtualServer{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// VirtualServerListWithCustomId Common 이 아닌 별도 Partition 으로 resource id 가상 서버 검색
func (g *GTM) VirtualServerListWithCustomId(id string) (*VirtualServer, error) {
	url := g.c.buildUrl(basePath, ServerResource, id, VirtualServerResource)
	resp := &VirtualServer{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/* Modify Resource */

func NewServerConfig() *Server {
	return new(Server)
}

func (r *Server) SetName(name string) *Server {
	r.Name = name
	return r
}

func (r *Server) SetPartition(partition string) *Server {
	r.Partition = partition
	return r
}

func (r *Server) SetDatacenter(datacenter string) *Server {
	if strings.Contains(datacenter, slash) {
		r.Datacenter = datacenter
	} else {
		r.Datacenter = fmt.Sprintf("/Common/%s", datacenter)
	}

	return r
}

func (r *Server) SetMonitor(monitor string) *Server {
	if strings.Contains(monitor, slash) {
		r.Monitor = monitor
	} else {
		r.Monitor = fmt.Sprintf("/Common/%s", monitor)
	}
	return r
}

func (r *Server) SetAddresses(ip ...string) *Server {
	for _, v := range ip {
		r.Addresses = append(r.Addresses, Address{
			Name: v,
		})
	}

	return r
}

func (r *Server) SetProduct(product string) *Server {
	r.Product = product
	return r
}

func (r *Server) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set Server.Name")
	}
	if r.Partition == "" {
		r.Partition = "Common"
	}
	if r.Datacenter == "" {
		return fmt.Errorf("please set Server.Datacenter")
	}

	return nil
}

func (g *GTM) CreateServer(serverConfig *Server) (*Server, error) {
	url := g.c.buildUrl(basePath, ServerResource)
	if err := serverConfig.Verify(); err != nil {
		return nil, newError(400, "server value verify fail: "+err.Error())
	}

	resp := &Server{}
	body, err := json.Marshal(serverConfig)
	if err != nil {
		return nil, newError(500, "CreateServer.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(HTTPPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DeleteServer(id string) (*Server, error) {
	url := g.c.buildUrl(basePath, ServerResource, CommonId(id))
	resp := &Server{}
	err := g.c.iControlRequest(HTTPDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewVirtualServerConfig() *VirtualServer {
	return new(VirtualServer)
}

func (r *VirtualServer) SetName(name string) *VirtualServer {
	r.Name = name
	return r
}

func (r *VirtualServer) SetAddress(ip, port string) *VirtualServer {
	r.Destination = ip
	if port == "" {
		r.Destination += ":0"
	} else {
		r.Destination += port
	}

	return r
}

func (r *VirtualServer) Verify() error {
	if r.Name == "" {
		return fmt.Errorf("please set VirtualServer.Name")
	}
	if r.Destination == "" {
		return fmt.Errorf("please set VirtualServer.Destination")
	}

	return nil
}

func (g *GTM) AddVirtualServer(serverId string, virtualServerConfig *VirtualServer) (*VirtualServer, error) {
	url := g.c.buildUrl(basePath, ServerResource, CommonId(serverId), VirtualServerResource)
	if err := virtualServerConfig.Verify(); err != nil {
		return nil, newError(400, "virtualServer value verify fail: "+err.Error())
	}

	resp := &VirtualServer{}
	body, err := json.Marshal(virtualServerConfig)
	if err != nil {
		return nil, newError(500, "CreateVirtualServer.Marshal fail: "+err.Error())
	}
	err = g.c.iControlRequest(HTTPPost, url, body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) RemoveVirtualServer(serverId, virtualServerId string) (*VirtualServer, error) {
	url := g.c.buildUrl(basePath, ServerResource, CommonId(serverId), VirtualServerResource, virtualServerId)
	resp := &VirtualServer{}
	err := g.c.iControlRequest(HTTPDelete, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
