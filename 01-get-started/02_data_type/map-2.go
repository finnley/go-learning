package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	mapLearn()

	//goplLearn()
}

func goplLearn() {
	fmt.Println("======================= gopl =======================")
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages["bob"]++

	fmt.Println("======================= for range =======================")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Println("======================= sort1 =======================")
	sort1(ages)

	fmt.Println("======================= sort2 =======================")
	map1()

	fmt.Println("======================= dedup =======================")
	// 面的dedup程序读取多行输入，但是只打印第一次出现的行。dedup程序通过map来表示所有的输入行所对应的set集合，以确保已经在集合存在的行不会被重复打印。
	dedup()
}

func dedup() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		} else {
			fmt.Println("already exist")
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func sort1(ages map[string]int) {
	//var names []string
	// 因为我们一开始就知道names的最终大小，因此给slice分配一个合适的大小将会更有效。下面的代码创建了一个空的slice，但是slice的容量刚好可以放下map中全部的key：
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	fmt.Printf("%v\n", names) // [alice charlie bob]
	sort.Strings(names)
	fmt.Printf("%v\n", names) // [alice bob charlie]

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func map1() {
	// map类型的零值是nil，也就是没有引用任何哈希表
	var ages map[string]int
	fmt.Println(ages == nil)    // true
	fmt.Println(len(ages) == 0) // true
	// map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常：
	//ages["carol"] = 21 // 这个有问题，

}

func mapLearn() {
	fmt.Println("======================= 定义 =======================")
	// go 语言中的map的key和value类型申明时就要指明
	// 1. 字面值
	m1 := map[string]string{
		"m1": "v1",
	}
	fmt.Printf("%v\n", m1)

	// 2. make 函数创建map，创建时里面不添加元素
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Printf("%v\n", m2)

	// 3. 定义一个空 map
	m3 := map[string]string{}
	fmt.Printf("%v\n", m3)

	// map 中的key不是所有类型都支持，该类型需要支持 == 或者 != 操作， 比如 int rune
	// 切片就不支持,切片不支持等于和不等于操作,但是切片可以和nil比较
	//a := []int{1, 2, 3}
	//b := []int{1, 2, 3}
	//if a == b {
	//
	//}
	//var m1 map[[]int]string

	// 数组支持
	//a := [3]int{1, 2, 3}
	//b := [3]int{1, 2, 3}
	//var m1 map[[3]int]int

	fmt.Println("======================= 操作 =======================")
	// map 的基本操作
	m := map[string]string{
		"a": "va",
		"b": "vb",
		"d": "",
	}
	// 1. 增加，修改
	m["c"] = "vc"
	m["b"] = "vb1"
	fmt.Printf("%v\n", m) // map[a:va b:vb1 c:vc d:]

	// 2. 查询
	v, ok := m["d"]
	fmt.Println(v, ok) //  true

	// 3. 删除
	//delete(m, "a")
	//delete(m, "e")
	//delete(m, "a")
	//fmt.Printf("%v\n", m)

	// 4. 遍历
	// Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。
	// 在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。这是故意的，每次都使用随机的遍历顺序可以强制要求程序不会依赖具体的哈希函数实现
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("================ 按顺序遍历key/value对 ===========")
}

/**
禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
*/
