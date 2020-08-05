package router

import (
	"back/handler/companyIntroHandler"
	"github.com/kataras/iris/v12/core/router"
)

// 路由分组
func Route(v1 router.Party){
	
	
	
	// 前台
	frontGroup := v1.Party("/intro")
	{

		frontGroup.Get("", companyIntroHandler.GetI)

	}
		//roadGroup := frontGroup.Party("/road")
		//{
		//	//roadGroup.Get(":id", companyRoadHandler.GetId)
		//	//获取公司历程
		//	roadGroup.Get("", companyRoadHandler.GetI)
		//}
		//bannerGroup := frontGroup.Party("/banner")
		//{
		//	// 获取轮播图
		//	bannerGroup.Get("",homeBannerHandler.GetI)
		//}
		//categoryGroup := frontGroup.Party("/category")
		//{
		//	categoryGroup.Get("",productCategoryHandler.GetI)
		//}
		//detailGroup := frontGroup.Party("/detail")
		//{
		//	detailGroup.Get("",productDetailHandler.GetI)
		//}


	//// 后台登录入口
	//v1.Post("/login", handler.CodeLogin)
	//// swagger:route /auth
	//backGroup := v1.Party("/auth",middleware.AuthMiddleware())
	backGroup := v1.Party("/auth")
	{
		// swagger:route Get /intro
		introGroup := backGroup.Party("/intro")
		{

			introGroup.Get("", companyIntroHandler.GetAll)
		}
	}
	//		introGroup.Post("", companyIntroHandler.Create)
	//		// 更新数据
	//		introGroup.Put(":id", companyIntroHandler.Update)
	//		// 删除单个
	//		//introGroup.Delete(":id", companyIntroHandler.Delete)
	//		// 批量删除
	//		introGroup.Delete("", companyIntroHandler.Deletes)
	//
	//	}
	//	roadGroup := backGroup.Party("/road")
	//	{
	//		roadGroup.Get("", companyRoadHandler.GetAll)
	//		roadGroup.Post("", companyRoadHandler.Create)
	//		roadGroup.Put(":id", companyRoadHandler.Update)
	//		//roadGroup.Delete(":id", companyRoadHandler.Delete) 方法冗余，不使用
	//		roadGroup.Delete("", companyRoadHandler.Deletes)
	//	}
	//	bannerGroup := backGroup.Party("/banner")
	//	{
	//
	//		bannerGroup.Get("", homeBannerHandler.GetAll)
	//		bannerGroup.Post("", homeBannerHandler.Create)
	//		bannerGroup.Put(":id", homeBannerHandler.Update) //测试标记
	//		//bannerGroup.Delete(":id", homeBannerHandler.Delete)
	//		bannerGroup.Delete("", homeBannerHandler.Deletes)
	//		// 更新图片，上传图片
	//		bannerGroup.Post("/upload", homeBannerHandler.UploadBanner)
	//	}
	//	categoryGroup := backGroup.Party("/category")
	//	{
	//		categoryGroup.Get("", productCategoryHandler.GetAll)
	//		categoryGroup.Post("", productCategoryHandler.Create)
	//		categoryGroup.Put(":id", productCategoryHandler.Update)
	//		//categoryGroup.Delete(":id", productCategoryHandler.Delete)
	//		categoryGroup.Delete("", productCategoryHandler.Deletes)
	//	}
	//	detailGroup := backGroup.Party("/detail")
	//	{
	//
	//		detailGroup.Get("", productDetailHandler.GetAll)
	//		detailGroup.Post("", productDetailHandler.Create)
	//		detailGroup.Put(":id", productDetailHandler.Update)
	//		//detailGroup.Delete(":id", productDetailHandler.Delete)
	//		detailGroup.Delete("", productDetailHandler.Deletes)
	//		detailGroup.Post("/upload", productDetailHandler.UploadDetail)
	//	}
	//}

}
