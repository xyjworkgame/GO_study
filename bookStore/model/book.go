/*
@Time : 2020/5/21 22:32
@Author : Firewine
@File : book
@Software: GoLand
@Description:
*/
package model
//Book 结构体
type Book struct {
	ID int
	Title string
	Author string
	Price float64
	Sales int
	Stock int
	ImgPath string
}