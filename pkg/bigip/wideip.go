package bigip

type WideIpList struct {
	Kind  string   `json:"kind"`
	Items []WideIp `json:"items"`
}

type WideIp struct {
	Kind                 string `json:"kind"`
	Name                 string `json:"name"`
	Partition            string `json:"partition"`
	FullPath             string `json:"fullPath"`
	Generation           int    `json:"generation"`
	Enabled              bool   `json:"enabled"`
	FailureRcode         string `json:"failureRcode"`
	FailureRcodeResponse string `json:"failureRcodeResponse"`
	FailureRcodeTTL      int    `json:"failureRcodeTtl"`
	LastResortPool       string `json:"lastResortPool"`
	MinimalResponse      string `json:"minimalResponse"`
	PersistCidrIpv4      int    `json:"persistCidrIpv4"`
	PersistCidrIpv6      int    `json:"persistCidrIpv6"`
	Persistence          string `json:"persistence"`
	PoolLbMode           string `json:"poolLbMode"`
	TTLPersistence       int    `json:"ttlPersistence"`
	Pools                []struct {
		Name          string `json:"name"`
		Partition     string `json:"partition"`
		Order         int    `json:"order"`
		Ratio         int    `json:"ratio"`
		NameReference struct {
			Link string `json:"link"`
		} `json:"nameReference"`
	} `json:"pools"`
	PoolsCname []struct {
		Name          string `json:"name"`
		Partition     string `json:"partition"`
		Order         int    `json:"order"`
		Ratio         int    `json:"ratio"`
		NameReference struct {
			Link string `json:"link"`
		} `json:"nameReference"`
	} `json:"poolsCname"`
}

func (c *client) WideIpList(ResourceType string) (*WideIpList, error) {
	url := c.buildUrl(basePath, WideIpResource, ResourceType)
	resp := &WideIpList{
		Items: make([]WideIp, 0),
	}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) WideIp(id, ResourceType string) (*WideIp, error) {
	url := c.buildUrl(basePath, WideIpResource, ResourceType, CommonId(id))
	resp := &WideIp{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) WideIpWithCustomId(id, ResourceType string) (*WideIp, error) {
	url := c.buildUrl(basePath, WideIpResource, ResourceType, id)
	resp := &WideIp{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
