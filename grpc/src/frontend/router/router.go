package router

import (
	"context"
	"github.com/kataras/iris/v12"
	"rpc/grpc/src/frontend/driver/rpc_connect"
	"rpc/grpc/src/frontend/handler"
)
var (
	// rpc 全文连接
	Ctx context.Context
)

func Router(router iris.Party) {
	Ctx = context.Background()


	ProductServiceServer := rpc_connect.MustConnGRPC("ProductServiceServer")
	productsHandler := handler.NewProductsHandler(ProductServiceServer,Ctx)




	apiV1 := router.Party("/api/v1")
	{
		products := apiV1.Party("/products")
		{
			products.Get("/",productsHandler.Gets)
			products.Get("/:int",productsHandler.GetById)
		}
	}

	//router.Post("/login",) 登录
	manageV1 := router.Party("/manage/v1")
	{
		products := manageV1.Party("/product")
		{
			products.Get("/list.do",productsHandler.GetsDo)
			products.Get("/search.do",productsHandler.GetSearchDo)
			products.Get("/detail.do",productsHandler.GetById)
			products.Put("/set_sale_status.do",productsHandler.PutStatusDo)
			products.Post("/save.do",productsHandler.PostProductDo)
		}
		//manageV1.Post("/upload.do",handler.U)  图片上传,分别为富文本图片，和正常图片

	}
}
