/*
@Time : 2020/5/17 22:14
@Author : Firewine
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"html/template"
	"net/http"
)

//创建处理器函数
func testTemplate(w http.ResponseWriter, r *http.Request) {

	//t,_ := template.ParseFiles("E:\\GO\\GoWeb\\src\\web02_template\\index.html")
	//通过must函数让GO去自动处理异常
	t := template.Must(template.ParseFiles("E:\\GO\\GoWeb\\src\\web02_template\\index.html",
		"E:\\GO\\GoWeb\\src\\web02_template\\index2.html"))

	//t.Execute(w,"我是后台传输数据")
	//响应数据在index2文件显示
	t.ExecuteTemplate(w,"index2.html","我在2显示")

}
func main() {

	//创建多路复用器
	//mux := http.NewServeMux()
	//http.HandleFunc("/", handler)
	http.HandleFunc("/testTemplate", testTemplate)


	//创建路由
	//http.ListenAndServe(":8080",nil)
	http.ListenAndServe(":8080",nil)
}