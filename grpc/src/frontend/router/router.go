package router

import (
	"github.com/kataras/iris/v12"
	"rpc/grpc/src/frontend/handler"
)

func Router(router iris.Party) {

	apiV1 := router.Party("/api/v1")
	{
		products := apiV1.Party("/products")
		{
			products.Get("/",handler.GetProducts)
		}
	}
}
