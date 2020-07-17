/*
@Time : 2020/5/21 23:20
@Author : Firewine
@File : bookDao
@Software: GoLand
@Description:
*/
package dao

import (
	"LearnGO/src/bookStore/model"
	"LearnGO/src/bookStore/utils"
)

// getBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books, nil
}