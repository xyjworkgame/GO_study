package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	vipee := viper.New()
	vipee.SetConfigFile("config.yaml") //指定配置文件
	vipee.AddConfigPath("/Viper/config/")//配置文件路径  不起作用

	if err := vipee.ReadInConfig();err != nil{// 读取配置文件
		panic(fmt.Errorf("fatal error config file: %s \n",err))
	}

	vipee.WatchConfig()  //监控配置文件
	vipee.OnConfigChange(func(in fsnotify.Event) {
		//配置文件发生变更之后会调用的回调函数
		fmt.Println("config file changed:",in.Name)
	})


}
