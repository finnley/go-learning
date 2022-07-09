package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var b = "Go语言"
	for index, runeValue := range b {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	/**
	  U+0047 'G' starts at byte position 0
	  U+006F 'o' starts at byte position 1
	  U+8BED '语' starts at byte position 2
	  U+8A00 '言' starts at byte position 5
	 */
	fmt.Println()

	const nihongo = "Go语言"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runValue, width := utf8.DecodeLastRuneInString(nihongo[i:])
		fmt.Printf("width: %d\n", width)
		fmt.Printf("%#U starts at byte position %d\n", runValue, i)
		w = width
	}
	/**
	U+8A00 '言' starts at byte position 0
	U+8A00 '言' starts at byte position 3
	U+FFFD '�' starts at byte position 6
	U+FFFD '�' starts at byte position 7
	 */
}
