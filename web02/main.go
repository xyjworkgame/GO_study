/*
@Time : 2020/5/17 10:44
@Author : Firewine
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {

}

func (m *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//fmt.Fprintln(writer,"通过自己创建的处理器处理请求")
	fmt.Fprintln(writer,"通过自己创建详细配置的server的处理器处理请求")
}


func main() {

	myHandler := MyHandler{}
	//http.Handle("/myHandler",&myHandler)
	//创建server结构，详细配置字段
	server := http.Server{
		Addr: ":8080",
		Handler: &myHandler,
		ReadHeaderTimeout: 2*time.Second,
	}
	//http.ListenAndServe(":8080",nil)
	//导入自定义的Server
	server.ListenAndServe()
}
