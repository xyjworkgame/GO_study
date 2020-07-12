/*
@Time : 2020/5/16 23:44
@Author : Firewine
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"net/http"
)


//创建处理器函数
func handler(w http.ResponseWriter, r*http.Request) {
	//fmt.Fprintf(w,"hello world!",r.URL.Path)
	fmt.Fprintln(w,"通过自己的多路复用器请求")
}
func main() {

	//创建多路复用器
	mux := http.NewServeMux()
	//http.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)


	//创建路由
	//http.ListenAndServe(":8080",nil)
	http.ListenAndServe(":8080",mux)
}