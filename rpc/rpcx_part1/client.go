package main

import (
	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

func main() {
	// #1定义了使用什么方式来实现服务发现。 在这里我们使用最简单的,点对点
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	// #2  创建了 XClient， 并且传进去了 FailMode、 SelectMode 和默认选项。 FailMode 告诉客户端如何处理调用失败：重试、快速返回，或者 尝试另一台服务器。 SelectMode 告诉客户端如何在有多台服务器提供了同一服务的情况下选择服务器。
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3 定义了请求：这里我们想获得 10 * 20 的结果。 当然我们可以自己算出结果是 200，但是我们仍然想确认这与服务器的返回结果是否一致
	args := &example.Args{
		A: 10,
		B: 20,
	}

	// #4定义了响应对象， 默认值是0值， 事实上 rpcx 会通过它来知晓返回结果的类型，然后把结果反序列化到这个对象
	reply := &example.Reply{}

	// #5调用了远程服务并且同步获取结果。
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
