package model

import (
	"github.com/jinzhu/gorm"
)

type CompanyIntro struct {
	gorm.Model
	Title   string `gorm:"column:title;type:varchar(255);comment:'title'" json:"title"`
	Content string `gorm:"column:content;type:text;comment:'内容'" json:"content"`
	Status  int    `gorm:"column:status;default:1;comment:'状态(2:禁用;1:正常)'" json:"status"`
	SeqNo   int    `gorm:"column:seq_no;comment:'排列顺序'" json:"seq_no"`
}
type CompanyIntroJson struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	SeqNo   int    `json:"seq_no"`
}
