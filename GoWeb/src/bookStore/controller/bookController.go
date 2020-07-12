/*
@Time : 2020/5/22 0:02
@Author : Firewine
@File : bookController
@Software: GoLand
@Description:
*/
package controller

import (
	"GoWeb/src/bookStore/dao"
	"html/template"
	"net/http"
)

func GetBooks(w http.ResponseWriter,r *http.Request){
	//调用dao方法
	books , _ := dao.GetBooks()

	t := template.Must(template.ParseFiles("./src/bookStore/views/pages/manager/book_manager.html"))

	t.Execute(w,books)

}