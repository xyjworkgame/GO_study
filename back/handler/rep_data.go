package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"log"
	"net/http"
)

func Fault(c iris.Context, errno int, msg string) {
	log.Println(msg)
}
// 返回多种类型数据
func RespData(c iris.Context,options context.JSON) {


	c.JSON(http.StatusOK, options)
}
// 返回data
func Success(c iris.Context, options context.JSON) {

	c.JSON(http.StatusOK,options)

}


