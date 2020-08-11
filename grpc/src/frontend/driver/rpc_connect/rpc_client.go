package rpc_connect

import (
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"
	"time"
)

var (
	zkAddr   = flag.String("zkAddr", "192.168.20.190:2181", "zookeeper address")
	basePath = flag.String("base", "/mall", "prefix path")
)

// 生成客户端session，
func MustConnGRPC(servicePath string) client.XClient {

	d := client.NewZookeeperDiscovery(*basePath, servicePath, []string{*zkAddr}, nil)

	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	option.RPCPath = share.DefaultRPCPath
	option.ConnectTimeout = 10 * time.Second
	option.SerializeType = protocol.MsgPack
	option.CompressType = protocol.None
	option.BackupLatency = 10 * time.Millisecond
	//option.Group = "1"  这次切换为随机

	xclient := client.NewXClient(servicePath, client.Failfast, client.RandomSelect, d, option)

	return xclient
}
