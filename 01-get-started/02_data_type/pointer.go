package main

import "fmt"

func main0206() {
	a := 10
	//fmt.Println(&a) // 0xc00001c0b0

	var p *int = &a
	fmt.Println(p)  // 0xc00001c0b0
	fmt.Println(&a) // 0xc00001c0b0

	// 通通过指针间接修改变量值
	*p = 132
	fmt.Println(a) // 132

	// 注意指针可以指向变量地址，但是不能指向常量地址
	//const MAX int = 100
	//fmt.Println(&MAX) // cannot take the address of MAX
}

func main() {
	// 定义指针默认值为nil，nil可以理解为指向内存地址编号为0的空间，内存0-255内存空间为系统占用，不允许用户读写操作
	// 定义了指针，但没有指向任何内存空间，然后又修改了这个内存空间的值
	//var p *int // panic: runtime error: invalid memory address or nil pointer dereference
	//*p = 123
	//fmt.Println(*p)

	// 开辟数据类型大小的空间，返回值为指针类型，也就是通过new创建空间，并把这个空间的地址赋给一个指针
	// 就相当于开了个房，然后把房的钥匙拿到手，就可以通过这个钥匙进入到房间
	// new会动态的分配内存空间，不需要关心内存释放问题，GO的GC会自动回收
	//new(数据类型)
	var p *int
	p = new(int)
	*p = 123
	fmt.Println(*p) // 123
}
