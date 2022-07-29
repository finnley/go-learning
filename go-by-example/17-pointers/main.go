package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("07-form-validate-pointers:", &i)
}

/**
initial: 1
zeroval: 1
zeroptr: 0
07-form-validate-pointers: 0xc00000a0b8


允许在程序中通过 引用传递 来传递值和数据结构。

我们将通过两个函数：zeroval 和 zeroptr 来比较 指针 和 值。 zeroval 有一个 int 型参数，所以使用值传递。 zeroval 将从调用它的那个函数中得到一个实参的拷贝：ival。

zeroptr 有一个和上面不同的参数：*int，这意味着它使用了 int 指针。 紧接着，函数体内的 *iptr 会 解引用 这个指针，从它的内存地址得到这个地址当前对应的值。 对解引用的指针赋值，会改变这个指针引用的真实地址的值。


*/
