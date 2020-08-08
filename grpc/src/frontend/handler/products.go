package handler

import "github.com/kataras/iris/v12"

// 前端
type productsHandler struct {}

func NewCompanyIntroHandler() *productsHandler {
	return &productsHandler{

	}
}



// 产品搜索及动态排序
func (p *productsHandler)GetProducts(ctx iris.Context) {
	ctx.URLParam("categoryId")
	ctx.URLParam("keyword")
	ctx.URLParamDefault("pageNum","1")
	ctx.URLParamDefault("pageSize","20")
	ctx.URLParamDefault("orderBy","")//price_desc ,price_asc
}