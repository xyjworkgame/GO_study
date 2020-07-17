/*
@Time : 2020/5/20 19:21
@Author : Firewine
@File : userController
@Software: GoLand
@Description:
*/
package controller

import (
	"LearnGO/src/bookStore/dao"
	"html/template"
	"net/http"
)

// Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("username")

	//调用userdao中验证用户名和密码
	user,_ := dao.CheckUserNameAndPassword(username,password)
	if user != nil && user.ID != 0{
		//用户名和密码正确
		t := template.Must(template.ParseFiles("./src/bookStore/views/pages/user/login_success.html"))
		t.Execute(w,"")
	}else {
		//用户名密码不正确
		t := template.Must(template.ParseFiles("./src/bookStore/views/pages/user/login.html"))
		t.Execute(w,"用户名或密码不正确")
	}
}

//注册
func Regist(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("username")
    email := r.PostFormValue("email")
	//调用userdao中验证用户名和密码
	user,_ := dao.CheckUserName(username)
	if user.ID> 0{
		//用户名存在
		t := template.Must(template.ParseFiles("./src/bookStore/views/pages/user/regist.html"))
		t.Execute(w,"用户名已存在")
	}else {
		//用户名存在，保存用户
		dao.SaveUser(username,password,email)
		t := template.Must(template.ParseFiles("./src/bookStore/views/pages/user/regist_success.html"))
		t.Execute(w,"")
	}
}
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")

	//调用userdao中验证用户名和密码
	user,_ := dao.CheckUserName(username)
	if user.ID > 0{
		//用户名存在
		w.Write([]byte("用户名已存在"))
	}else {
		//用户名不存在
		w.Write([]byte("用户名可用"))
	}

}
