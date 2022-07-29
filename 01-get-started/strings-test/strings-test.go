package main

import (
	"fmt"
	"strings"
)

func main1() {
	hostName := "server-server-udp1"
	serverId := hostName
	if strings.HasPrefix(hostName, "server-") {
		//fmt.Printf("split: %#v", strings.Split(hostName, "server-"))
		fmt.Printf("split: %#v", hostName[7:])
		hostName = strings.Split(hostName, "server-")[1]
	} else {
		serverId = fmt.Sprintf("server-%s", hostName)
	}
	fmt.Println(hostName, serverId)

}

func main() {
	s := "hello world hello world"
	//str := "wo"

	//以str为分隔符，将s切分成多个子串，结果中**不包含**str本身。如果str为空则将s切分成Unicode字符列表。
	//如果s中没有str子串，则将整个s作为[]string的第一个元素返回。
	//参数n表示最多切分出几个子串，超出的部分将不再切分，最后一个n包含了所有剩下的不切分。
	//如果n为0，则返回nil；如果n小于0，则不限制切分个数，全部切分
	index := strings.SplitN(s, " ", 2)
	fmt.Println(len(index)) //2
	fmt.Println(index)
}
