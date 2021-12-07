package bigip

type PoolList struct {
	Kind  string `json:"kind"`
	Items []Pool `json:"items"`
}

type Pool struct {
	Kind                     string `json:"kind"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	SelfLink                 string `json:"selfLink"`
	AlternateMode            string `json:"alternateMode"`
	DynamicRatio             string `json:"dynamicRatio"`
	Enabled                  bool   `json:"enabled"`
	FallbackMode             string `json:"fallbackMode"`
	LoadBalancingMode        string `json:"loadBalancingMode"`
	ManualResume             string `json:"manualResume"`
	QosHitRatio              int    `json:"qosHitRatio"`
	QosHops                  int    `json:"qosHops"`
	QosKilobytesSecond       int    `json:"qosKilobytesSecond"`
	QosLcs                   int    `json:"qosLcs"`
	QosPacketRate            int    `json:"qosPacketRate"`
	QosRtt                   int    `json:"qosRtt"`
	QosTopology              int    `json:"qosTopology"`
	QosVsCapacity            int    `json:"qosVsCapacity"`
	QosVsScore               int    `json:"qosVsScore"`
	TTL                      int    `json:"ttl"`
	VerifyMemberAvailability string `json:"verifyMemberAvailability"`
	MembersReference         struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"membersReference"`
}

type PoolMemberList struct {
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selfLink"`
	Items    []PoolMember `json:"items"`
}

type PoolMember struct {
	Kind                      string `json:"kind"`
	Name                      string `json:"name"`
	Partition                 string `json:"partition"`
	FullPath                  string `json:"fullPath"`
	Generation                int    `json:"generation"`
	SelfLink                  string `json:"selfLink"`
	Enabled                   bool   `json:"enabled"`
	LimitMaxBps               int    `json:"limitMaxBps"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus"`
	LimitMaxConnections       int    `json:"limitMaxConnections"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus"`
	LimitMaxPps               int    `json:"limitMaxPps"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus"`
	MemberOrder               int    `json:"memberOrder"`
	Monitor                   string `json:"monitor"`
	Ratio                     int    `json:"ratio"`
}

func (c *client) PoolList(ResourceType string) (*PoolList, error) {
	url := c.buildUrl(basePath, PoolResource, ResourceType)
	resp := &PoolList{
		Items: make([]Pool, 0),
	}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) Pool(id, ResourceType string) (*Pool, error) {
	url := c.buildUrl(basePath, PoolResource, ResourceType, CommonId(id))
	resp := &Pool{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) PoolWithCustomId(id, ResourceType string) (*Pool, error) {
	url := c.buildUrl(basePath, PoolResource, ResourceType, id)
	resp := &Pool{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) PoolMemberList(id, ResourceType string) (*PoolMemberList, error) {
	url := c.buildUrl(basePath, PoolResource, ResourceType, CommonId(id), MemberResource)
	resp := &PoolMemberList{
		Items: make([]PoolMember, 0),
	}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) PoolMemberListWithCustomId(id, ResourceType string) (*PoolMemberList, error) {
	url := c.buildUrl(basePath, PoolResource, ResourceType, id, MemberResource)
	resp := &PoolMemberList{
		Items: make([]PoolMember, 0),
	}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
