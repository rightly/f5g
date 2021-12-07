package bigip

type DatacenterList struct {
	Kind  string       `json:"kind"`
	Items []Datacenter `json:"items"`
}

type Datacenter struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Enabled    bool   `json:"enabled"`
	Location   string `json:"location,omitempty"`
}

func (g *GTM) DatacenterList() (*DatacenterList, error) {
	url := g.c.buildUrl(basePath, DatacenterResource)
	resp := &DatacenterList{
		Items: make([]Datacenter, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) Datacenter(id string) (*Datacenter, error) {
	url := g.c.buildUrl(basePath, DatacenterResource, CommonId(id))
	resp := &Datacenter{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) DatacenterWithCustomId(id string) (*Datacenter, error) {
	url := g.c.buildUrl(basePath, DatacenterResource, id)
	resp := &Datacenter{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
