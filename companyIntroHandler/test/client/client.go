
package main

import (
	"context"
	"flag"
	example "github.com/rpcxio/rpcx-examples"
	"log"
	"time"


	"github.com/smallnest/rpcx/client"
)

var (
	zkAddr   = flag.String("zkAddr", "192.168.60.100:2181", "zookeeper address")
	basePath = flag.String("base", "/youpin/services", "prefix path")
)

func main() {
	flag.Parse()

	d := client.NewZookeeperDiscovery(*basePath, "getI", []string{*zkAddr}, nil)
	xclient := client.NewXClient("getI", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {

		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}

}