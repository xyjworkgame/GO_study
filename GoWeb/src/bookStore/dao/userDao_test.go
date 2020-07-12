/*
@Time : 2020/5/19 22:29
@Author : Firewine
@File : userDao_test.go
@Software: GoLand
@Description:
*/
package dao

import (
	"fmt"
	"testing"
)

func TestSaveUser(t *testing.T) {
	fmt.Println("开始测试dao中方法testSaveUser")
	err := SaveUser("admin","123456","aa@qq.com")
	fmt.Println(err)
}

func TestCheckUserName(t *testing.T) {
	fmt.Println("开始测试dao中方法testCheckUserName")
	user,_ := CheckUserName("admin")
	fmt.Println(user.Email)
}

func TestCheckUserNameAndPassword(t *testing.T) {
	fmt.Println("开始测试dao中方法testCheckUserNameAndPassword")
	user,_ := CheckUserNameAndPassword("admin","123456")
	fmt.Println(user)
}
