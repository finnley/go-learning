package main

import (
	"fmt"
	"strings"
)

func main() {
	// 常见操作
	// 1. 是否包含某个子串
	var name string = "china:\"大中国\""
	fmt.Println(strings.Contains(name, "大中国")) // true

	// 2. 子串在字符串中的位置
	fmt.Println(strings.Index(name, "大中国")) // 7

	// 3. 统计出现的次数
	fmt.Println(strings.Count(name, "\"")) // 2

	// 4. 前缀和后缀
	fmt.Println(strings.HasPrefix(name, "c")) // true
	fmt.Println(strings.HasSuffix(name, "\"")) // true

	// 5. 大小写转换
	fmt.Println(strings.ToUpper("china")) // CHINA
	fmt.Println(strings.ToLower("CHINA")) // china

	// 6. 字符串的比较 字符串的比较是 ascii 的比较 ，小于返回-1,大于返回1，相等返回0
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("b", "a")) // 1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("ab", "a")) // 1 逐个比较

	// 7.去除空格 和指定字符串
	fmt.Println(strings.TrimSpace(" china ")) // china
	fmt.Println(strings.TrimLeft("chinca", "c")) // hinca 从左边开始满足要求的字符
	fmt.Println(strings.Trim("chinac", "c")) // hina
	fmt.Println(strings.Trim("chinach", "ch")) // ina
	fmt.Println(strings.Trim("chinac", "ch")) // ina
	fmt.Println(strings.Trim("chinah", "ch")) // ina

	// 8.分割
	fmt.Println(strings.Split("china shanghai", " ")) // [china shanghai]

	// 9.合并
	arrs := strings.Split("china shanghai", " ")
	fmt.Println(strings.Join(arrs, "-")) // china-shanghai

	// 10.字符串替换
	fmt.Println(strings.Replace("china:1949:china", "china", "China", 1)) // China:1949:china
	fmt.Println(strings.Replace("china:1949:china", "china", "China", 2)) // China:1949:China

	// 11.判断字符串s是否包含chars字符串中的任一字符
	fmt.Println(strings.ContainsAny(name, "r")) // false
	fmt.Println(strings.ContainsAny(name, "a")) // true

	// 12.判断字符串s是否包含符文书r
	fmt.Println(strings.ContainsRune(name, 104))
	// 13.将字符串s以空白符分割，返回一个切片
	fmt.Println(strings.Fields("china i love you")) // [china i love you]

}
