/*
@Time : 2020/5/18 23:38
@Author : Firewine
@File : db
@Software: GoLand
@Description:
*/
package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init(){
	Db, err = sql.Open("mysql","root:123456@tcp(localhost:3306)/bookstore")
	if err != nil{
		panic(err)
	}
}