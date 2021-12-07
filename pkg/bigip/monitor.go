package bigip

type MonitorList struct {
	Kind     string    `json:"kind"`
	SelfLink string    `json:"selfLink"`
	Items    []Monitor `json:"items"`
}

type Monitor struct {
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	SelfLink           string `json:"selfLink"`
	DefaultsFrom       string `json:"defaultsFrom,omitempty"`
	Destination        string `json:"destination"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	ProbeTimeout       int    `json:"probeTimeout"`
	Reverse            string `json:"reverse"`
	Timeout            int    `json:"timeout"`
	Transparent        string `json:"transparent"`
}

func (g *GTM) MonitorList(MonitorType string) (*MonitorList, error) {
	url := g.c.buildUrl(basePath, MonitorResource, MonitorType)
	resp := &MonitorList{
		Items: make([]Monitor, 0),
	}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) Monitor(id, MonitorType string) (*Monitor, error) {
	url := g.c.buildUrl(basePath, MonitorResource, MonitorType, CommonId(id))
	resp := &Monitor{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GTM) MonitorWithCustomId(id, MonitorType string) (*Monitor, error) {
	url := g.c.buildUrl(basePath, MonitorResource, MonitorType, id)
	resp := &Monitor{}
	err := g.c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
