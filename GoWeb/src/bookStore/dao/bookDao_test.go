/*
@Time : 2020/5/22 0:10
@Author : Firewine
@File : bookDao_test.go
@Software: GoLand
@Description:
*/
package dao

import (
	"GoWeb/src/bookstore/dao"
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books ,_ := dao.GetBooks()

	fmt.Println(books)
}

