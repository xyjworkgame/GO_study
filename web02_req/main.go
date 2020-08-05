/*
@Time : 2020/5/17 19:17
@Author : Firewine
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"LearnGO/src/web01_db/model"
	json2 "encoding/json"
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))
	// //获取请求体重内容的长度
	// len := r.ContentLength
	// //创建byte切片
	// body := make([]byte, len)
	// //将请求体中的内容读到body中
	// r.Body.Read(body)
	// //在浏览器中显示请求体中的内容
	// fmt.Fprintln(w, "请求体中的内容有：", string(body))

	//解析表单，在调用r.Form之前必须执行该操作
	// r.ParseForm()
	//获取请求参数
	//如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	//那么参数值都可以得到，并且form表单中的参数值在ULR的参数值的前面
	//fmt.Fprintln(w, "请求参数有：", r.Form)
	//fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)
	//通过直接调用FormValue方法和PostFormValue方法直接获取请求参数的值
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是：", r.PostFormValue("username"))
}
func testJsonRes(w http.ResponseWriter, r *http.Request) {

	//设置响应类型
	w.Header().Set("Content-Type", "application/json")
	//创建User
	user := model.User{
		ID:       1,
		Username: "adimin",
		Password: "222",
		Email:    "34343",
	}
	//将struct 转换为json
	json, _ := json2.Marshal(user)
	//将json格式数据返回客户端
	w.Write(json)
}
func testRedire(w http.ResponseWriter, r *http.Request) {

	//设置重定向，location为地址
	w.Header().Set("location", "https://www.baidu.com")
	//	设置响应的状态码
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedire", testRedire)
	http.ListenAndServe(":8080", nil)
}
