package bigip

type TopologyRecordList struct {
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selfLink"`
	Items    []TopologyRecord `json:"items"`
}

type TopologyRecord struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	SelfLink   string `json:"selfLink"`
	Order      int    `json:"order"`
	Score      int    `json:"score"`
}

func (c *client) TopologyRecordList() (*TopologyRecordList, error) {
	url := c.buildUrl(basePath, TopologyResource)
	resp := &TopologyRecordList{
		Items: make([]TopologyRecord, 0),
	}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) TopologyRecord(id string) (*TopologyRecord, error) {
	url := c.buildUrl(basePath, TopologyResource, urlEncoding(id))
	resp := &TopologyRecord{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
