package bigip

const (
	basePath 	= "/mgmt/tm/gtm"

	WideIpResource 		= "wideip"
	MonitorResource 	= "monitor"
	PoolResource 		= "pool"
	ServerResource 		= "server"
	TopologyResource	= "topology"
	RegionResource		= "region"
	DatacenterResource	= "datacenter"
	VirtualServerResource = "virtual-servers"

	AType		= "a"
	AAAAType	= "aaaa"
	CNAMEType	= "cname"
	MXType		= "mx"
	NAPTRType	= "naptr"
	SRVType		= "srv"

	// Pool LB Mode
	RoundRobinMode 	= "round-robin"
	GlobalAvailabilityMode = "global-availability"
	RandomMode 		= "random"
	RatioMode 		= "ratio"
	TopologyMode	= "topology"

	HTTPGet 	= "GET"
	HTTPPost	= "POST"
	HTTPPut		= "PUT"

	// Server Product
	GenericHostProduct = "generic-host"
)