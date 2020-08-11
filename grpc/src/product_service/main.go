package main

import (
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"rpc/grpc/src/product_service/driver/mysql"
	"rpc/grpc/src/product_service/service"
	"time"
)

var (
	addr2    = flag.String("addr2", "localhost:8972", "server address")
	addr1    = flag.String("addr1", "localhost:8971", "server address")
	zkAddr   = flag.String("zkAddr", "192.168.20.190:2181", "zookeeper address")
	basePath = flag.String("base", "/mall", "prefix path")
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	db := mysql.DB
	defer func() {
		_ = db.Close()
	}()
	flag.Parse()



	go createServer2(addr2, "group=1")
	go createServer1(addr1, "group=2")
	select {}

}
func createServer1(addr *string, meta string) {
	s := server.NewServer()

	addRegistryPlugin(s, *addr)

	if err := s.Register(new(service.ProductServiceServer), meta); err != nil {
		return
	}
	if err := s.Serve("tcp", *addr); err != nil {
		return
	}
}

func createServer2(addr *string, meta string) {
	s := server.NewServer()

	addRegistryPlugin(s, *addr)

	if err := s.Register(new(service.ProductServiceServer), meta); err != nil {
		return
	}
	if err := s.Serve("tcp", *addr); err != nil {
		return
	}
}

func addRegistryPlugin(s *server.Server, addr string) {

	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
