/*
@Time : 2020/5/17 13:46
@Author : Firewine
@File : db.go
@Software: GoLand
*/
package utils

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var(
	Db *sql.DB
	err error
)


func init() {
	Db,err = sql.Open("mysql","root:123456@tcp(localhost:3306)/test")
	if err != nil{
		panic(err.Error())
	}

}