package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}
type Rect struct {
}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}
func main() {
	rpc.Register(new(Rect))
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {

			fmt.Println("new client: ",conn.RemoteAddr())
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}