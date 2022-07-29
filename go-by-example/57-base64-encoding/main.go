package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // YWJjMTIzIT8kKiYoKSctPUB+

	// 解码可能会返回错误，如果不确定输入信息格式是否正确， 那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec)) // abc123!?$*&()'-=@~
	fmt.Println()

	// 使用 URL base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc) // YWJjMTIzIT8kKiYoKSctPUB-
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec)) // abc123!?$*&()'-=@~
}

/**
标准 base64 编码和 URL base64 编码的 编码字符串存在稍许不同（后缀为 + 和 -）， 但是两者都可以正确解码为原始字符串。
 */