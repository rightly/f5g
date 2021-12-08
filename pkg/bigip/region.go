package bigip

type RegionList struct {
	Kind  string   `json:"kind,omitempty"`
	Items []Region `json:"items,omitempty"`
}

type Region struct {
	Kind          string `json:"kind,omitempty"`
	Name          string `json:"name"`
	Partition     string `json:"partition"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	RegionMembers []struct {
		Name string `json:"name,omitempty"`
	} `json:"regionMembers,omitempty"`
}

func (g *GTM) RegionList() (*RegionList, error) {
	url := g.c.buildUrl(basePath, regionResource)
	resp := &RegionList{
		Items: make([]Region, 0),
	}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) Region(id string) (*Region, error) {
	url := g.c.buildUrl(basePath, regionResource, CommonId(id))
	resp := &Region{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) RegionWithCustomId(id string) (*Region, error) {
	url := g.c.buildUrl(basePath, regionResource, id)
	resp := &Region{}
	err := g.c.iControlRequest(httpGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
