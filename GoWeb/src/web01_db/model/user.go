package model

import (
	"GoWeb/src/web01_db/utils"
	"fmt"
)

//User 结构体
type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

//AddUser 添加User的方法一
func (user *User) AddUser() error {
	// 写sql 语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常:", err)
		return err
	}
	// 执行
	_, err2 := inStmt.Exec("admin1", "1234", "admin@11.com")
	if err2 != nil {
		fmt.Println("执行出现异常:", err2)
		return err2
	}
	return nil
}

//AddUser 添加User的方法二
func (user *User) AddUser2() error {
	// 写sql 语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 执行
	_, err2 := utils.Db.Exec(sqlStr, "admin2", "123456", "admin@11.com")
	if err2 != nil {
		fmt.Println("执行出现异常:", err2)
		return err2
	}
	return nil
}

// GetUserById 根据用户的id从数据库中查询一条记录
func (user *User) GetUserById() (*User, error) {
	// sql
	sqlStr := "select id,username,password,email from users where id = ?"
	// exec
	row := utils.Db.QueryRow(sqlStr, user.ID)
	// 声明
	var id int64
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

// GetUsers 获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	//sql
	sqlStr := "select id,username,password,email from users"
	//	执行
	rows,err := utils.Db.Query(sqlStr)
	if err != nil{
		return nil, err
	}
	// 创建切片
	var users []*User
	for rows.Next(){
		var id int64
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users,nil
}
