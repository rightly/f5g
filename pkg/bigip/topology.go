package bigip

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

func (g *GTM) TopologyRecordList() (*TopologyRecordList, error) {
	url := g.c.buildUrl(basePath, TopologyResource)
	resp := &TopologyRecordList{
		Items: make([]TopologyRecord, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) TopologyRecord(id string) (*TopologyRecord, error) {
	url := g.c.buildUrl(basePath, TopologyResource, IdEncoding(id))
	resp := &TopologyRecord{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
