package main

import "fmt"

func main() {
	//var courses6 = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	//subCourse6 := courses6[1:4] // [scrapy tornado python]
	//subCourse62 := subCourse6[1:4]
	//fmt.Printf("%T, %v\n", subCourse62, subCourse62) // []string, [tornado python]
	//subCourse621 := subCourse6[1:]
	//fmt.Printf("%v\n", subCourse621) // [tornado python]
	//subCourse622 := subCourse6[1:4]
	//fmt.Printf("%v\n", subCourse622) // [tornado python asyncio]
	//fmt.Printf("%v\n", courses6)     // [django scrapy tornado python asyncio]
	//subCourse623 := subCourse6[1:5]
	//fmt.Printf("%v\n", subCourse623) // panic: runtime error: slice bounds out of range [:5] with capacity 4

	var courses = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	subCourse := courses[1:4]                        // [scrapy tornado python]
	subCourse_1 := subCourse[1:3]                    // 没有超出 subCourse 的长度
	fmt.Printf("%T, %v\n", subCourse_1, subCourse_1) // []string, [tornado python]
	subCourse_2 := subCourse[1:4]                    // 结束位置超出了subCourse 的长度
	fmt.Printf("%T, %v\n", subCourse_2, subCourse_2) // []string, [tornado python asyncio]
	//subCourse_3 := subCourse[1:5]                    // 结束位置超出了 courses 的长度, panic
	//fmt.Printf("%T, %v\n", subCourse_3, subCourse_3) // error

	subCourse_4 := subCourse[1:]
	fmt.Printf("%v\n", subCourse_4) // [tornado python]
	fmt.Println(subCourse_4[:0])
}
