# f5g

F5 BIGIP Rest Client

## Resources
- Server
- Pool
- Wide IP
- Monitor
- Datacenter
- Topology
- Region

## Usage
### GTM Client
```go
// create client
client := bigip.New(scheme, domain, user, pw)
client.DisableTLSVerify()
gtm := client.GTM()

...use gtm
```
### Server
#### get server
```go
// server list
serverList, err := gtm.ServerList()
if err != nil {
    // error handling
}

// server
server, err := gtm.Server("server_10.10.10.10")
if err != nil {
    // error handling
}
```
#### get virtual server
```go
// virtual server list
vserverList, err := gtm.VirtualServerList()
if err != nil {
    // error handling
}

// virtual server
// use (serverId, virtualServerId string)
vserverList, err := gtm.VirtualServer("server_10.10.10.10", "server_10.10.10.10")
if err != nil {
    // error handling
}
```
#### create server 
```go
// set server config
server := bigip.NewServerConfig()
server.SetName("server_10.10.10.10").
        SetAddresses("10.10.10.10").
        SetProduct(bigip.ServerGenericHost).
        SetDatacenter("KR-1").
        SetMonitor("gateway_icmp").
        SetPartition("Common")

// server create
serverResp, err := gtm.CreateServer(server)
if err != nil {
    // error handling
}

```
#### add virtual server
```go
// set virtual server config
virtserver := bigip.NewVirtualServerConfig()
virtserver.SetName("server_10.10.10.10").
	    SetAddress("10.10.10.10", "")

// add virtual server
_, err = gtm.AddVirtualServer(serverResp.Name, virtserver)
if err != nil {
    // error handling
}
```
#### delete server
```go
// delete server
serverResp, err := gtm.DeleteServer("server_10.10.10.10")
if err != nil {
    // error handling
}
```
#### delete virtual server
```go
// remove virtual server
// use (serverId, virtualServerId string)
vserverResp, err := gtm.RemoveVirtualServer("server_10.10.10.10", "server_10.10.10.10")
if err != nil {
    // error handling
}
```
### Pool
#### get pool
```go
// pool list
poolList, err := gtm.PoolList(bigip.AType)
if e != nil {
    // error handling
}

// pool list
pool, err := gtm.Pool("pool-id", bigip.AType)
if e != nil {
// error handling
}
```
#### get pool member
```go
// pool member list
poolMemberList, err := gtm.PoolMemberList("pool-id", bigip.AType)
if e != nil {
    // error handling
}
```
#### create pool
```go
pool := bigip.NewPoolConfig()
pool.SetName("pool-id").
    SetMonitor("TCP_443").
    SetTTL(180)

_, err := gtm.CreatePool(pool, bigip.AType)
if err != nil {
    // error handling
}
```
#### add pool member
```go
poolMember := bigip.NewPoolMemberConfig()
poolMember.SetName("server_10.10.10.10")

_, err := gtm.AddPoolMember(poolMember, "pool-id", bigip.AType)
if err != nil {
    // error handling
}
```
#### delete pool
```go
_, err := gtm.DeletePool("pool-id", bigip.AType)
if err != nil {
    // error handling
}
```
#### remove pool member
```go
_, err := gtm.RemovePoolMember("pool-id", "server_10.10.10.10", bigip.AType)
if err != nil {
    // error handling
}
```