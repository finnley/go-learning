package main

import "fmt"

// SplitArray: 将数组元素平均分组
func SplitArray(array []int, part int) (result []interface{}) {
	// 1. 参数校验，数组长度为0 或者 划分组数大于数组元素个数
	if len(array) == 0 || len(array) < part {
		return nil
	}

	num := len(array) / part
	// 2.分组
	if len(array)%part == 0 {
		// a.元素数量刚好能够划分指定组数
		for i := 0; i < len(array); i += num {
			result = append(result, array[i:i+num])
		}
	} else {
		// b.最后一组不足组成一组，将最后一组中的每个元素追加到其他组中
		// 先平均分组，最后不足一组的将每个元素一次分配到前面几组中
		// 比如有14个数组，分为3组，最后两个元素会追加到前两组中，各组数量为(5, 5, 4)

		// 每组元素个数
		elemNum := len(array) / part
		// 最后一个组元祖起始索引
		lastGroupIndex := len(array) - len(array)%part
		// 定义一个Map，每个字典元素为一组
		m := make(map[int][]int, elemNum)
		start := 0
		for i := 0; i < part; i++ {
			end := start + elemNum
			if start < lastGroupIndex {
				for j := start; j < end; j++ {
					m[i] = append(m[i], array[j])
				}
				start = end
			}
		}
		// 最后一组元素
		lastPart := array[lastGroupIndex:]
		// 最后不足一组的将每个元素一次分配到前面几组中
		for k, v := range lastPart {
			m[k] = append(m[k], v)
		}

		// 输出结果
		for _, v := range m {
			result = append(result, v)
		}
	}
	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := SplitArray(arr, 3)
	fmt.Println(result) // [[1 2 3 10] [4 5 6] [7 8 9]]

	//abc := make([]int, 2)
	//abc = append(abc, 1, 2, 3)
	//fmt.Println(abc)

	//s := S{}
	//p := &s
	////f(s)
	////g(s)
	////f(p)
	//g(p)
	slice := make([]int, 0, 1)
	slice = append(slice, 1)
	slice1 := appendTest(slice)
	fmt.Println(slice1)
}

func appendTest(slice []int) (result []int) {
	result = append(slice, 10)
	fmt.Println(result)
	return
}

type S struct {
}

func f(x interface{}) {
}
func g(x *interface{}) {
}
