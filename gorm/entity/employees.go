package entity

import "github.com/jinzhu/gorm"

type TmpEmployee struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);comment:'name'" json:"name"`
	Password string `gorm:"column:password;type:varchar(255);comment:'password'" json:"password"`
	Age      int    `gorm:"column:age;type:int(3);comment:'age'" json:"age"`
	SeqNo	int `gorm:"column:seq_no;type:int(11);comment:'sort'"`
}
