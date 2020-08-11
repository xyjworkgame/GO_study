package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"rpc/grpc/src/product_service/model"
	"time"
)

var (
	DB *gorm.DB
)

const (

	// 开发数据库连接
	devConn190 = "root:123456-abc@tcp(192.168.20.190)/mall?charset=utf8&parseTime=True&loc=Local"
	devConn    = "root:123456@tcp(192.168.60.100)/okzm?charset=utf8&parseTime=True&loc=Local"
)

func init() {
	var err error
	// 通过环境变量ZM_DEBUG选择数据库连接

	DB, err = gorm.Open("mysql", devConn190)
	if err != nil {
		log.Fatal(err)
	}
	DB.LogMode(true)
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "mall_" + defaultTableName;
	}
	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	DB.SingularTable(true)
	DB.DB().SetConnMaxLifetime(2 * time.Hour)

	// 数据库迁移
	DB.AutoMigrate(new(model.Product))
}
