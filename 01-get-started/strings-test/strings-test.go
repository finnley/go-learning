package main

import (
	"fmt"
	"strings"
)

func main() {
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
