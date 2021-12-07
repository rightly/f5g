package bigip

const (
	blank = " "
	slash = "/"
	tilde = "~"
)

const (
	basePath = "/mgmt/tm/gtm"

	// Resource
	WideIpResource        = "wideip"
	MonitorResource       = "monitor"
	PoolResource          = "pool"
	MemberResource        = "members"
	ServerResource        = "server"
	TopologyResource      = "topology"
	RegionResource        = "region"
	DatacenterResource    = "datacenter"
	VirtualServerResource = "virtual-servers"

	// RRType
	AType     = "a"
	AAAAType  = "aaaa"
	CNAMEType = "cname"
	MXType    = "mx"
	NAPTRType = "naptr"
	SRVType   = "srv"

	// Pool LB Mode
	RoundRobinMode         = "round-robin"
	GlobalAvailabilityMode = "global-availability"
	RandomMode             = "random"
	RatioMode              = "ratio"
	TopologyMode           = "topology"

	// API Method
	HTTPGet  = "GET"
	HTTPPost = "POST"
	HTTPPut  = "PUT"

	// Server Product
	ServerGenericHost = "generic-host"

	// Monitor Type
	MonitorHTTP        = "http"
	MonitorHTTPS       = "https"
	MonitorTCP         = "tcp"
	MonitorGatewayICMP = "gateway-icmp"
)