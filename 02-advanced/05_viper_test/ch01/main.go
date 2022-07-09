package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// 将配置文件映射成 struct
type ServerConfig struct {
	// 将配置文件中的 name 的值方位 ServiceName 中
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}

func main() {
	v := viper.New()
	// 文件路径如何设置?
	// v.SetConfigFile("config-debug.yaml") // 这种方式需要切换到main文件所在目录
	v.SetConfigFile("02-advanced/05_viper_test/ch01/config-debug.yaml") // 这种方式相对项目根目录，可以直接运行该文件
	// 读取文件，有可能出错，需要判断
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	name := v.Get("name")
	fmt.Println(name)

	// 映射成 struct
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig) // {use-web 8021}
	fmt.Printf("%s\n", v.Get("name"))

}
