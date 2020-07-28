package main

import (
	"context"
	"flag"
	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}
var (
	addr = flag.String("addr", "localhost:8972", "server address")
)
type Arith int
// 参数1 ： context 参数2 ： 包含了请求的数据A 和B ，参数3 ：指向reply结构体的指针
func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()


	s := server.NewServer()
	s.Register(new(example.Arith), "")

	s.Serve("tcp",*addr)
}