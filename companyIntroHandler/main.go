package main

import (
	mysql "companyIntroHandler/driver"
	repo "companyIntroHandler/repository"
	"flag"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	zkAddr   = flag.String("zkAddr", "192.168.20.190:2181", "zookeeper address")
	basePath = flag.String("base", "/home/companyIntro", "prefix path")
)

func main() {

	db := mysql.DB
	defer func() {
		_ = db.Close()
	}()

	repo.Intro = repo.NewCompanyIntro(db)
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	if err := s.RegisterFunctionName("GetI", "GetI", repo.GetI,"");err != nil{
		log.Println(err)
	}
	if err := s.RegisterFunctionName("GetAll", "GetAll",repo.GetAll, "");err != nil{
		log.Println(err)
	}
	//if err := s.RegisterName("GetI", new(repo.GetI), "");err != nil{
	//	log.Println(err)
	//}
	//if err := s.RegisterName("GetI", new(repo.GetI), "");err != nil{
	//	log.Println(err)
	//}
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
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
