package main

import (
	"GoWeb/src/bookStore/controller"
	"html/template"
	"net/http"
)

// IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板

	t := template.Must(template.ParseFiles("./src/bookStore/views/index.html"))
	//执行
	t.Execute(w, "")
}

func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./src/bookStore/views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./src/bookStore/views/pages"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/getBooks",controller.GetBooks)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	http.HandleFunc("/regist", controller.Regist)
	http.ListenAndServe(":8080", nil)
}
