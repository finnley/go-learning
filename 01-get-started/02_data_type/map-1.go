package main

import "fmt"

// map定义
// var 字典名 map[键类型]值类型
func main() {
	//使用make方式定义并初始化
	//var m map[int]string = make(map[int]string, 10)
	//m[1001] = "法师"
	//m[8888] = "兵哥"
	//fmt.Println(m) // map[1001:法师 8888:兵哥]

	// 使用简短声明方式定义并初始化
	m := map[int]string{1001: "法师", 8888: "兵哥", 1234: "星程", 3333: "哈哈"}
	fmt.Println(m) // map[1001:法师 1234:星程 3333:哈哈 8888:兵哥]
	fmt.Println(m[12345])
	fmt.Printf("%#v\n", m[12345])

	//// 需要注意的是map里面的数据是无序的
	//// k: 键 v: 值
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
}

func main0211() {
	// // 键不能是数组 键必须支持 == 或者 != 操作，不能是切片 map 或函数
	//m := make(map[[]int]string)

	// make定义完可以直接添加
	m := make(map[int]string)
	m[1001] = "法师"
	m[1005] = "兵哥"

	// 通过key判断value是否存在
	if v, ok := m[1001]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	fmt.Println(len(m))
	// 删除map中的数据
	delete(m, 1001)
	fmt.Println(m) // map[1005:兵哥]
	fmt.Println(len(m))
}
