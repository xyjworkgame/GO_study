package companyIntroHandler

import (
	example "back/handler/companyIntroHandler/proto"
	"context"
	"encoding/json"
	"flag"
	"github.com/kataras/iris/v12"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	zkAddr   = flag.String("zkAddr", "192.168.20.190:2181", "zookeeper address")
	basePath = flag.String("base", "/home/companyIntro", "prefix path")
	xclient  client.XClient
)


func GetI(ctx iris.Context) {
	flag.Parse()
	d := client.NewZookeeperDiscovery(*basePath, "GetI", []string{*zkAddr}, nil)

	xclient = client.NewXClient("GetI", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	args := &example.RequestGetI{}
	reply := &example.ResponseGetI{}
	err := xclient.Call(context.Background(), "GetI", args, reply)
	if err != nil{
		log.Println("failed to call: %v",err)
		xclient.Close()
	}
	defer xclient.Close()
	log.Println(reply.String())
	marshal, err := json.Marshal(reply.String())
	//log.Println(reply.CompanyIntro,marshal)
	log.Println(marshal)
	ctx.JSON(reply.String())

}
func GetAll(ctx iris.Context) {
	flag.Parse()
	d := client.NewZookeeperDiscovery(*basePath, "GetAll", []string{*zkAddr}, nil)

	xclient = client.NewXClient("GetAll", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	args := &example.RequestGetAll{
		Status: 1,
		Page: &example.Page{
			PageNum: 1,
			PageSize: 20,
		},
	}
	reply := &example.ResponseGetAll{}
	err := xclient.Call(context.Background(), "GetAll", args, reply)
	if err != nil{
		log.Println("failed to call: %v",err)
		xclient.Close()
	}
	defer xclient.Close()
	log.Println(reply.String())
	marshal, err := json.Marshal(reply.String())
	//log.Println(reply.CompanyIntro,marshal)
	log.Println(marshal)
	ctx.JSON(reply.String())

}



