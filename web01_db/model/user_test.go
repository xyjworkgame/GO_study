/*
@Time : 2020/5/17 16:18
@Author : Firewine
@File : user_test.go
@Software: GoLand
@Description:
*/
package model

import (
	"fmt"
	"testing"
)
//TestMain函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始：")
	//通过m.Run()来执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	//t.Run("测试查询用户:", testUserGetUserById)
	////通过t.Run()来执行子测试函数
	//t.Run("测试添加用户:", testAddUser)
	t.Run("测试查询用户:", testGetUsers)

}

//如果函数名不是以Test开头，那么该函数默认不执行，我们可以将它设置成为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("子测试函数执行：")
	// user := &User{}
	//调用添加用户的方法
	// user.AddUser()
	// user.AddUser2()
}
// 测试获取一个User
func testUserGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := User{
		ID: 5,
	}
	//调用获取
	fmt.Println("kaishi")
	u,_ := user.GetUserById()
	fmt.Println(u)
}

//测试获取所有User
func testGetUsers(t *testing.T){
	fmt.Println("测试查询所有记录")
	user := &User{}
	//调用方法获取User的方法
	us,_ := user.GetUsers()
	//遍历切片
	for k,v := range us{
		fmt.Printf("第%v个用户是%v\n",k+1,v)
	}
}