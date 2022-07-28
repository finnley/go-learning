package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main2() {
	fmt.Println("golang连接redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.8.200:6379",
		Password: "123456",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func main() {
	fmt.Println("golang连接redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//添加键值对
	err = client.Set("golang", "yes", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("键golang设置成功")

	//通过键查询值
	val, err := client.Get("golang").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("键golang的值为: ", val)
	redis
}
