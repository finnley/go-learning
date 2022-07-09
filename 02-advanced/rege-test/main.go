package main

import (
	"fmt"
	"regexp"
)

func checkPasswordPolicy(password string, passwordPolicy string) error {
	if passwordPolicy == "" {
		return nil
	}
	if ok, err := regexp.MatchString(passwordPolicy, password); !ok {
		return fmt.Errorf("err: %s, %s", password, passwordPolicy)
	} else {
		return err
	}
}

func main() {
	//fmt.Println("(?![0-9A-Z1+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{16,18}$")
	//fmt.Println(checkPasswordPolicy("N9B8NheB#vlZiW&b", "(?![0-9A-Z1+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{16,18}$"))
	//fmt.Println(checkPasswordPolicy("12345678901234567", "(?![0-9A-Z]+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{16,18}$"))
	//fmt.Println(checkPasswordPolicy("qwertyuiopasdfghj", "(?![0-9A-Z]+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{16,18}$"))
	//fmt.Println(checkPasswordPolicy("QWERTYUIOPASDFGHJKL", "(?![0-9A-Z]+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{16,18}$"))
	//fmt.Println(checkPasswordPolicy("aA9B8NheB#vlZiW&%_", "(?![0-9A-Z]+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{16,18}$"))

	//match, _ := regexp.MatchString("(?![0-9A-Z]+$)(?![0-9a-z]+$)(?![a-zA-Z]+$)[0-9A-Za-z~!]{16,18}$", "aA9B8NheBsfdfdf!~")
	//match, _ := regexp.MatchString("(?![0-9]+$)[A-Za-z!]$", "sd4df")
	//fmt.Println(match) //true

	//r, _ := regexp.Compile(`f([a-z]+)`)
	// go get -u github.com/dlclark/regexp2
	r, err := regexp.Compile("(?![0-9A-Z]+$)[0-9A-Za-z~!@#$%!^*+-=_&]{4,10}$")

	//r, err := regexp2.Compile("[0-9A-Za-z~!@#$%!^*+-=_&]{4,6}$", 0)
	//r, err := regexp2.Compile("[0-9A-Za-z~!@#$%!^*+-=_&]{4,6}$", 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.MatchString("1233343545")) //true
	}

}
