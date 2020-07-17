package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml") //指定配置文件
	viper.AddConfigPath("/Viper/config/")//配置文件路径

	if err := viper.ReadInConfig();err != nil{// 读取配置文件
		panic(fmt.Errorf("fatal error config file: %s \n",err))
	}

	viper.WatchConfig()  //监控配置文件
	viper.OnConfigChange(func(in fsnotify.Event) {
		//配置文件发生变更之后会调用的回调函数
		fmt.Println("config file changed:",in.Name)
	})


}
