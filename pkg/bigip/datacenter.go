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

func (c *client) DatacenterList() (*DatacenterList, error) {
	url := c.buildUrl(basePath, DatacenterResource)
	resp := &DatacenterList{
		Items: make([]Datacenter, 0),
	}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) Datacenter(id string) (*Datacenter, error) {
	url := c.buildUrl(basePath, DatacenterResource, CommonId(id))
	resp := &Datacenter{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) DatacenterWithCustomId(id string) (*Datacenter, error) {
	url := c.buildUrl(basePath, DatacenterResource, id)
	resp := &Datacenter{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
