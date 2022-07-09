package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//s := "hello, world"
	//fmt.Println(len(s)) // 12

	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}

	fmt.Println("================================================")

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

/**
0	H	1
1	e	1
2	l	1
3	l	1
4	o	1
5	,	1
6	 	1
7	世	3
10	界	3

每一次调用DecodeRuneInString函数都返回一个r和长度，r对应字符本身，长度对应r采用UTF8编码后的编码字节数目

0	'H'	72
1	'e'	101
2	'l'	108
3	'l'	108
4	'o'	111
5	','	44
6	' '	32
7	'世'	19990
10	'界'	30028

*/
