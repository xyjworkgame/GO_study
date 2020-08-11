package handler

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/smallnest/rpcx/client"
	"log"
	demo "rpc/grpc/pb"
)

// 前端
type productsHandler struct {
	xclient client.XClient
	ctx context.Context
}

func NewProductsHandler(xclient client.XClient,ctx context.Context) *productsHandler {
	return &productsHandler{
		xclient: xclient,
		ctx: ctx,
	}
}

// 产品搜索及动态排序
func (p *productsHandler)Gets(ctx iris.Context) {
	ctx.URLParam("categoryId")
	ctx.URLParam("keyword")
	ctx.URLParamDefault("pageNum","1")
	ctx.URLParamDefault("pageSize","20")
	ctx.URLParamDefault("orderBy","")//price_desc ,price_asc
}
// 产品详细数据
func (p *productsHandler) GetById(ctx iris.Context) {
	productId := ctx.Params().GetInt64Default("id", 0)
	req := &demo.GetProductByIdReq{
		ProductId: productId,
	}
	res := &demo.GetProductByIdRes{}

	err := p.xclient.Call(p.ctx, "GeById", req, res)

	if err != nil {
		//log.Fatalf("failed to call: %v", err)
		log.Panicf("failed to call: %v", err)
	}
	fmt.Println(res.Status)
}

func (p *productsHandler) GetsDo(ctx iris.Context) {
	ctx.URLParamDefault("pageNum","1")
	ctx.URLParamDefault("pageSize","20")
}

func (p *productsHandler) GetSearchDo(ctx iris.Context) {
	ctx.URLParamDefault("productName","")
	ctx.URLParamDefault("productId","")
	// 这里判断 上面两个 是否都为空
	ctx.URLParamDefault("pageNum","1")
	ctx.URLParamDefault("pageSize","20")
}

func (p *productsHandler) PutStatusDo(ctx iris.Context) {
	ctx.URLParamDefault("productName","")
	ctx.URLParamDefault("status","")
}

func (p *productsHandler) PostProductDo(ctx iris.Context) {
	//ctx.ReadJSON()
}