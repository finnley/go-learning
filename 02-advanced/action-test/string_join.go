package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"
	str := "STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,ONLY_FULL_GROUP_BY,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"
	fmt.Println(str)
	arr := strings.Split(str, ",")
	fmt.Println(arr)
	s := make([]string, 0)
	for _, val := range arr {
		if val != "ONLY_FULL_GROUP_BY" {
			s = append(s, val)
		}
	}
	str = strings.Join(s, ",")
	fmt.Println(str)
}
