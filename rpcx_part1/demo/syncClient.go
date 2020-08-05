package main


// 服务发现
/*
Peer to Peer: 客户端直连每个服务节点。 the client connects the single service directly. It acts like the client type.
Peer to Multiple: 客户端可以连接多个服务。服务可以被编程式配置。
Zookeeper: 通过 zookeeper 寻找服务。
Etcd: 通过 etcd 寻找服务。
Consul: 通过 consul 寻找服务。
mDNS: 通过 mDNS 寻找服务（支持本地服务发现）。
In process: 在同一进程寻找服务。客户端通过进程调用服务，不走TCP或UDP，方便调试使用。
*/

import (
	"context"
	"flag"
	"log"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr2 = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}