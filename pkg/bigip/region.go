package bigip

type RegionList struct {
	Kind  string   `json:"kind"`
	Items []Region `json:"items"`
}

type Region struct {
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	Partition     string `json:"partition"`
	FullPath      string `json:"fullPath"`
	Generation    int    `json:"generation"`
	SelfLink      string `json:"selfLink"`
	RegionMembers []struct {
		Name string `json:"name"`
	} `json:"regionMembers"`
}

func (g *GTM) RegionList() (*RegionList, error) {
	url := g.c.buildUrl(basePath, RegionResource)
	resp := &RegionList{
		Items: make([]Region, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) Region(id string) (*Region, error) {
	url := g.c.buildUrl(basePath, RegionResource, CommonId(id))
	resp := &Region{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) RegionWithCustomId(id string) (*Region, error) {
	url := g.c.buildUrl(basePath, RegionResource, id)
	resp := &Region{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
