package bigip

type ServerList struct {
	Kind     string   `json:"kind"`
	Items    []Server `json:"items"`
}

type Server struct {
	Kind              string        `json:"kind"`
	Name              string        `json:"name"`
	Partition         string        `json:"partition"`
	FullPath          string        `json:"fullPath"`
	Generation        int           `json:"generation"`
	Datacenter        string        `json:"datacenter"`
	Enabled           bool          `json:"enabled"`
	Monitor           string        `json:"monitor,omitempty"`
	Product           string        `json:"product"`
	Addresses         []Address     `json:"addresses"`
}

type Address struct {
	Name        string `json:"name"`
	DeviceName  string `json:"deviceName"`
	Translation string `json:"translation"`
}


type VirtualServer struct {
	Kind     string `json:"kind"`
	Items    []struct {
		Kind                      string `json:"kind"`
		Name                      string `json:"name"`
		FullPath                  string `json:"fullPath"`
		Generation                int    `json:"generation"`
		Destination               string `json:"destination"`
		Enabled                   bool   `json:"enabled"`
		TranslationAddress        string `json:"translationAddress"`
		TranslationPort           int    `json:"translationPort"`
	} `json:"items"`
}

func (c *client) ServerList() (*ServerList, error) {
	url := c.buildUrl(basePath, ServerResource)
	resp := &ServerList{
		Items:    make([]Server, 0),
	}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) Server(id string) (*Server, error) {
	url := c.buildUrl(basePath, ServerResource, CommonId(id))
	resp := &Server{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) ServerWithCustomId(id string) (*Server, error) {
	url := c.buildUrl(basePath, ServerResource, id)
	resp := &Server{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) VirtualServer(id string) (*VirtualServer, error) {
	url := c.buildUrl(basePath, ServerResource, CommonId(id), VirtualServerResource)
	resp := &VirtualServer{}
	err := c.iControlRequest(HTTPGet, url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}