package main

import (
	"flag"
	"github.com/kataras/iris/v12"
	"log"
	"rpc/grpc/src/frontend/router"
)

func main() {

	flag.Parse()
	app := iris.Default()

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 跨域请求 设置

	//rpc_connect.MustConnGRPC("ProductServiceServer")

	root := app.Party("/")
	router.Router(root)



	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":8080"))
}
