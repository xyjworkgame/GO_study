package main

import (
	"back/router"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"log"
)

// api 处理使用 iris 来进行调用
// 只是用来处理前端处理json，
// 转化成 rpc 数据转发
func main() {

	app := iris.Default()

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},   //允许通过的主机名称
		AllowCredentials: true,
	})
	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions) // <- 对于预检很重要。

	// 使用路由分组
	router.Route(v1)

	if err := app.Run(iris.Addr(":8082"));err != nil{
		log.Fatal(err)
	}
}
