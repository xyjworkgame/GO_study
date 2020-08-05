package mysql

import (
	"companyIntroHandler/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

const (
	// 正式数据库连接
	prodConn = "user:password@tcp(host)/db?charset=utf8&parseTime=True&loc=Local"
	// 开发数据库连接
	devConn190 = "root:123456-abc@tcp(192.168.20.190)/okzm_cp?charset=utf8&parseTime=True&loc=Local"
	devConn = "root:123456@tcp(192.168.60.100)/okzm?charset=utf8&parseTime=True&loc=Local"
)

func init() {
	var err error
	// 通过环境变量ZM_DEBUG选择数据库连接
	if os.Getenv("ZM_DEBUG") == "off" {
		DB, err = gorm.Open("mysql", devConn190)
	} else {
		DB, err = gorm.Open("mysql", devConn190)
		DB.LogMode(true)
	}
	if err != nil {
		log.Fatal(err)
	}
	DB.DB().SetConnMaxLifetime(2 * time.Hour)


	// 数据库迁移
	DB.AutoMigrate(&model.CompanyIntro{})
}
