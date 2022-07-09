package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
)

func main() {
	//re := regexp2.MustCompile(`(?![0-9A-Z]+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{16,18}$`, 0)
	//if isMatch, err := re.MatchString(`FXGpg6mP1ZoBuuZts4dD`); isMatch {

	re := regexp2.MustCompile(`[0-9]{16,18}$`, 0)
	if isMatch, err := re.MatchString(`1234567890123456789`); isMatch {
		//do something
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ok: ", isMatch)
	} else {
		fmt.Println("failed: ", isMatch)
	}
}
