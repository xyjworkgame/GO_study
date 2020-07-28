package main

import (
	"GoWeb/gorm/entity"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/rand"
	"time"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@(192.168.60.100)/employee?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
	}
	Db.DB().SetConnMaxLifetime(2 * time.Hour)
	Db.AutoMigrate(entity.TmpEmployee{})

}

var ChanData chan entity.TmpEmployee

func main() {
	ChanData = make(chan entity.TmpEmployee, 10000)
	table := Db.Model(entity.TmpEmployee{})
	go CreateData()


	go func() {
		for {

			tmp := <-ChanData
			table.Create(&tmp)
			if len(ChanData) == 0{
				break
			}
		}
	}()
	go CreateData()
	go func() {
		for {

			tmp := <-ChanData
			table.Create(&tmp)
			if len(ChanData) == 0{
				break
			}
		}
	}()
	go func() {
		for {

			tmp := <-ChanData
			table.Create(&tmp)
			if len(ChanData) == 0{
				break
			}
		}
	}()
	for {

		tmp := <-ChanData
		table.Create(&tmp)

		if len(ChanData) == 0{
			break
		}
	}

	defer Db.Close()
}

func CreateData() {
	for i := 1; i < 500000; i++ {
		tmp := entity.TmpEmployee{
			Name:     RandomStr(200),
			Password: RandomStr(200),
			Age:      rand.Intn(100),
			SeqNo:    rand.Intn(5000000),
		}

		ChanData <- tmp
		if i%10000 == 0 {
			fmt.Printf("time : %s , finish count: %v \n",time.Now().Format("2006-01-02 15:04:05"),i)

		}

	}
}
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
