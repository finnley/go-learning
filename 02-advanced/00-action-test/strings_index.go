package main

import (
	"fmt"
	"strings"
)

const (
	host   = "https://support.actionsky.com/service_desk/browse/"
	key404 = "CR-276"
	urlTpl = "%v%v"
)

func main() {
	//创建和初始化字符串
	str1 := "Welcome to the online portal of nhooo"
	//str2 := "My dog name is Dollar"
	//str3 := "I like to play Ludo"    //显示字符串    fmt.Println("字符串 1: ", str1)
	//fmt.Println("字符串 2: ", str2)
	//fmt.Println("字符串 3: ", str3)
	res := strings.Index(str1, "the")
	fmt.Println(res)
	fmt.Println(str1[res+len("the "):])

	err := fmt.Errorf("%v processlist collect is enable but not to exist processlist file. Please check if there is a history in the directory, ", "2asdg")
	fmt.Println(err)
	// https://support.actionsky.com/service_desk/browse/CR-415

	s := fmt.Sprintf(urlTpl, host, key404)
	fmt.Println(s)
}
