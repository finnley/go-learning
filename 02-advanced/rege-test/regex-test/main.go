package main

import (
	"fmt"
	"regexp"
)

var IgnoredApiList = []string{
	"/user/login",
	"/user/logout",
	"/mysql/group/list",
	"/alert_record/list_unresolved_num",
}

func main() {
	isIgnored := false
	for i := range IgnoredApiList {
		regular := fmt.Sprintf("^(/v)([1-9]\\d|[1-9])(%s)(.*)", IgnoredApiList[i])
		reg := regexp.MustCompile(regular)
		isOk := reg.MatchString("/v3/mysql/group2/list")
		if isOk {
			fmt.Println("满足要求")
		}
		//if isOk, _ := regexp.MatchString(fmt.Sprintf("^(/v)([1-9]\\d|[1-9])(%s)(.*)", IgnoredApiList[i]), "/v3/user/logout"); isOk {
		//	fmt.Println("满足要求")
		//	//return true
		//}
		//if strings.HasPrefix("/v3/user/logout", fmt.Sprintf("^(/v)([1-9]\\d|[1-9])(%s)(.*)", IgnoredApiList[i])) {
		//	isIgnored = true
		//	//return
		//
		//}

	}
	fmt.Println(isIgnored)
}
