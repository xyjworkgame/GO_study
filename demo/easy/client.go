package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 传参,暂时改成小写，同一个包下冲突
type params struct {
	Width, Height int
}

func main() {
	//	1. 连接远程rpc 服务
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//	2. 调用方法
	//	面积
	ret := 0
	err2 := conn.Call("Rect.Area", params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(ret)

	//	周长
	err3 := conn.Call("Rect.Perimeter", params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(ret)
}
