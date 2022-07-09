package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

// 如何将线上和线下的配置文件隔离
// 不用修改任何代码而且线上和线下的配置文件能隔离开
// 通常在本地电脑上配置一个环境变量，比如设置一个 MXSHOP_DEBUG:true true表示开启debug,false表示不开启，然后使用viper读取环境变量

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

// viper 读取环境变量
// ~/.zshrc:
// export MXSHOP_DEBUG=true
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main1() {
	v := viper.New()
	v.SetConfigFile("02-advanced/05_viper_test/ch02/config-debug.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	fmt.Println(serverConfig)   //{use-web {127.0.0.1 3306}}
	fmt.Println(v.Get("mysql")) // map[host:127.0.0.1 port:3306]
}

// 配置环境隔离
func main2() {
	// 读取环境变量 设置环境变量如果想要生效，需要重启goland
	fmt.Println(GetEnvInfo("MXSHOP_DEBUG")) // true
	fmt.Println(GetEnvInfo("CAT_DEBUG"))    // false
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("02-advanced/05_viper_test/ch02/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("02-advanced/05_viper_test/ch02/%s-debug.yaml", configFilePrefix)
	}
	fmt.Println(configFileName) // 02-advanced/05_viper_test/ch02/config-debug.yaml
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)         //{use-web2 {127.0.0.1 3306}}
	fmt.Printf("%v\n", v.Get("name")) // use-web2
}

// 动态监控变化
func main() {
	// 读取环境变量 设置环境变量如果想要生效，需要重启goland
	fmt.Println(GetEnvInfo("MXSHOP_DEBUG")) // true
	fmt.Println(GetEnvInfo("CAT_DEBUG"))    // false
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("02-advanced/05_viper_test/ch02/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("02-advanced/05_viper_test/ch02/%s-debug.yaml", configFilePrefix)
	}
	fmt.Println(configFileName) // 02-advanced/05_viper_test/ch02/config-debug.yaml
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)         //{use-web2 {127.0.0.1 3306}}
	fmt.Printf("%v\n", v.Get("name")) // use-web2

	//05_viper_test 还可以动态监控变化
	//运行文件 修改配置文件，自动输出
	v.WatchConfig()
	// 如何知道变化的值 这种不会堵塞住，所以为了防止退出需要设置time.Sleep(time.Second * 300)
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})
	time.Sleep(time.Second * 300)
}
