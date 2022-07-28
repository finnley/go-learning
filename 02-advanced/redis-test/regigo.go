package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"strconv"
)

func stringReply(reply interface{}, err error) (string, error) {
	switch reply.(type) {
	case []byte, string, nil, redis.Error:
		return redis.String(reply, err)
	case int64:
		intger := reply.(int64)
		return strconv.Itoa(int(intger)), nil
	default:
		return fmt.Sprintf(`%s`, reply), nil
	}
}

type redisSlowLogItem struct {
	// 唯一标识符
	Id string `json:"id"`
	// 命令执行时间，格式：UNIX时间戳
	Time string `json:"time"`
	// 执行命令耗费时间，单位：微秒
	Duration int32 `json:"duration"`
	// 命令与命令参数
	Argv []string `json:"argv"`
	// 命令与命令参数属性
	Argc []string `json:"argc"`
	// 慢日志采集时间
	CollectTime string `json:"collect_time"`
}

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	defer conn.Close()

	reply, err := conn.Do("SLOWLOG", "get", "-1")
	if err != nil {
		fmt.Println("SLOWLOG get err=", err)
		return
	}
	fmt.Printf("%T\n", reply)
	for _, v := range reply.([]interface{}) {
		//fmt.Printf("%T\n", v)

		//var list []redisSlowLogItem
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			s := reflect.ValueOf(v)
			res, _ := stringReply(s, err)
			if err != nil {
				return
			}
			fmt.Printf("%v\n", res)

			//fmt.Println(s)
			for i := 0; i < s.Len(); i++ {
				e := s.Index(i)
				//fmt.Printf("%T\n", e)

				//fmt.Println(e.Interface().(int64))
				switch e.Interface().(type) {
				case int64:
					fmt.Println("222: ", e.Interface().(int64))
				case string:
					fmt.Println("333: ", e.Interface().(string))
				default:
					a, _ := redis.String(e.Interface(), err)
					fmt.Println("default: ", a)
				}
				//list = append(list, e.Interface().(redisSlowLogItem))
			}
		}
	}

	//var list []redisSlowLogItem
	//if reflect.TypeOf(reply).Kind() == reflect.Slice {
	//	fmt.Println(111)
	//	s := reflect.ValueOf(reply)
	//	fmt.Println(s)
	//	for i := 0; i < s.Len(); i++ {
	//		e := s.Index(i)
	//		//fmt.Println(e)
	//		fmt.Printf("%T\n", e)
	//		//fmt.Println(e.Interface().(redisSlowLogItem))
	//		//list = append(list, e.Interface().(redisSlowLogItem))
	//	}
	//}

	//fmt.Println(list)

	////fmt.Println(stringReply(reply, err))
	//ret, err := stringReply(reply, err)
	////fmt.Printf("%s\n", ret)
	//
	//infoMap, err := InfoParser(ret)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(infoMap)

}

type Res struct {
}
