# 注册中心


## 点对点 Peer2Peer
这里也不需要注册中心，监控，
rpcx  格式： network @ Host: port格式表示一项服务

## MultipleServers
固定的几台服务器提供相同的服务，采用这种模式，
可以用注册中心，也可以用编码的方式

## zookeeper

### 服务器
1. 如果只需要zookeeper特性，而不需要etcd，consul等特性，使用tag 来标识
2. 服务端使用Zookeeper唯一的工作就是设置ZooKeeperRegisterPlugin这个插件。
3. 它主要配置几个参数：
    * ServiceAddress: 本机的监听地址， 这个对外暴露的监听地址， 格式为tcp@ipaddress:port
    * ZooKeeperServers: Zookeeper集群的地址
    * BasePath: 服务前缀。 如果有多个项目同时使用zookeeper，避免命名冲突，可以设置这个参数，为当前的服务设置命名空间
    * Metrics: 用来更新服务的TPS
    * UpdateInterval: 服务的刷新间隔， 如果在一定间隔内(当前设为2 * UpdateInterval)没有刷新,服务就会从Zookeeper中删除
**插件必须在注册服务之前添加到Server中，否则插件没有办法获取注册的服务的信息**
服务端
```go
// go run -tags zookeeper server.go
func main() {
    flag.Parse()

    s := server.NewServer()
    addRegistryPlugin(s)

    s.RegisterName("Arith", new(example.Arith), "")
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
```
客户端
```go

// go run -tags zookeeper client.go
    d := client.NewZookeeperDiscovery(*basePath, "Arith",[]string{*zkAddr}, nil)
    xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
    defer xclient.Close()
```

## Etcd
构建一个高可用的分布式键值数据库
使用方式与zookeeper 非常相似，而且在go中很适用


### 服务器
插件 EtcdRegisterPlugin插件， 配置参数和Zookeeper的插件相同。
参数：
   * ServiceAddress: 本机的监听地址， 这个对外暴露的监听地址， 格式为tcp@ipaddress:port
   * EtcdServers: etcd集群的地址
   * BasePath: 服务前缀。 如果有多个项目同时使用zookeeper，避免命名冲突，可以设置这个参数，为当前的服务设置命名空间
   * Metrics: 用来更新服务的TPS
   * UpdateInterval: 服务的刷新间隔， 如果在一定间隔内(当前设为2 * UpdateInterval)没有刷新,服务就会从etcd中删除
   
## consul
实现分布式系统的服务发现与配置
支持分布式，高可用，可横向扩展
特性 ： 服务发现，通过DNS，http接口发现。 健康监测：。key/value ； 提供简单的http接口，结合其他工具实现动态配置、功能标记，选举。多数据中心： 
需要配置ConsulRegisterPlugin插件。
参数：
    * ServiceAddress: 本机的监听地址， 这个对外暴露的监听地址， 格式为tcp@ipaddress:port
    * ConsulServers: consul集群的地址
    * BasePath: 服务前缀。 如果有多个项目同时使用consul，避免命名冲突，可以设置这个参数，为当前的服务设置命名空间
    * Metrics: 用来更新服务的TPS
    * UpdateInterval: 服务的刷新间隔， 如果在一定间隔内(当前设为2 * UpdateInterval)没有刷新,服务就会从consul中删除
    
    
## mDNS 多播
mDNS主要实现了在没有传统DNS服务器的情况下使局域网内的主机实现相互发现和通信，使用的端口为5353，遵从dns协议，使用现有的DNS信息结构、名语法和资源记录类型。并且没有指定新的操作代码或响应代码。

## Inprocess
用于进程内的测试 在开发过程中，可能不能直接连接线上的服务器直接测试，而是写一些mock程序作为服务，这个时候就可以使用这个registry, 测试通过在部署的时候再换成相应的其它registry
```go
func main() {
    flag.Parse()

    s := server.NewServer()
    addRegistryPlugin(s)

    s.RegisterName("Arith", new(example.Arith), "")

    go func() {
        s.Serve("tcp", *addr)
    }()

    d := client.NewInprocessDiscovery()
    xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
    defer xclient.Close()

    args := &example.Args{
        A: 10,
        B: 20,
    }

    for i := 0; i < 100; i++ {

        reply := &example.Reply{}
        err := xclient.Call(context.Background(), "Mul", args, reply)
        if err != nil {
            log.Fatalf("failed to call: %v", err)
        }

        log.Printf("%d * %d = %d", args.A, args.B, reply.C)

    }
}

func addRegistryPlugin(s *server.Server) {

    r := client.InprocessClient
    s.Plugins.Add(r)
}

```

