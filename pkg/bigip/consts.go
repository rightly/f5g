package bigip

const (
	blank = " "
	slash = "/"
	tilde = "~"
)

// GTM
const (
	basePath = "/mgmt/tm/gtm"

	// Resource
	wideIpResource        = "wideip"
	monitorResource       = "monitor"
	poolResource          = "pool"
	memberResource        = "members"
	serverResource        = "server"
	topologyResource      = "topology"
	regionResource        = "region"
	datacenterResource    = "datacenter"
	virtualServerResource = "virtual-servers"
	aWideIpResource       = "pools"
	cnameWideIpResource   = "pools-cname"

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

	// Pool FB Mode
	ReturnToDNS = "return-to-dns"

	// API Method
	httpGet    = "GET"
	httpPost   = "POST"
	httpPut    = "PUT"
	httpDelete = "DELETE"

	// Server Product
	ServerGenericHost = "generic-host"

	// Monitor Type
	MonitorHTTP        = "http"
	MonitorHTTPS       = "https"
	MonitorTCP         = "tcp"
	MonitorGatewayICMP = "gateway-icmp"
)
