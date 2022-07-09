# get started

## main 函数要点
无参数、无返回值
main 方法必须要在 main 包里面
go run 01_len_of_strings.go 就可以执行
如果文件不叫 01_len_of_strings.go，则需要 go build 得到可执行文件，然后直接运行可执行文件即可

## 包声明
语法形式：package xxx
字母和下划线组合
可以和文件夹不同名字
可以一个文件夹下的声明一致
比如有一个same的文件夹，里面有a.go 和 b.go 文件，它们的包名都是same，如果修改其中一个文件的包名为 not_sname，就会出现编译都不会通过，会出现下面提示：
Multiple packages in the directory: same, not_same

## 包引入
语法形式：import [alias] xxx
如果一个包引入了但是没有使用，会报错
匿名引入：前面多一个下划线，比如：_ "github.com/jinzhu/gorm/dialects/sqlite"
为什么需要匿名引用？因为包里面会有包初始化的方法，匿名引用就是为了执行里面的 init() 方法的，就没有其他效果了，比如使用sql时通常是匿名引入一个driver
除了匿名引用，Golang 会帮你自动引入你代码里面用得包，在你不用之后会删除

## string声明
双引号引起来，则内容双引号需要使用\转义，比如 println("He said:\" Hello World \" ")
`反引号引起来，则内部的 ` 需要\转义
string 长度很特殊
    字节长度：和编码无关，用 len(str) 获取
    字符数量：和编码有关，用编码库来计算
字符串拼接直接使用+，注意有些语言支持string和别的类型拼接，但是golang不可以
常用的strings操作：
    查找和替换
    大小写转换
    子字符串相关
    相等

## rune
直观理解，就是字符
rune 不是 byte
rune 本质是 int32，一个rune 有4个字节
rune 在很多语言里面是没有的，与之对应的 golang 没有char 类型，rune 不是数字，也不是char，也不是 byte
实际中不太常用

## 基本数据类型
bool: true, false
int8, int16, int32, int64, int
uint8, uint16, uint32, uint64, uint
float32, float64

## byte
字节，本质是uint8

类型总结
golang数字类型明确标注了长度，有无符号
golang不会帮我们做类型转换，类型不同无法通过编译。因此string 只能和 string 拼接
golang 有一个很特殊的类型 rune 类型，接近一般语言的 char 或者 character 的概念，非面试情况下，可以理解为 rune=字符

## 变量声明
首字母的大小写控制了变量的可访问性

## 常量

## 数组

## 切片
切片操作是有限的，不支持随机增删（即没有add, delete）方法，需要自己写代码
切片只有append操作
切片支持子切片操作，和原本切片共享底层数组
共享底层：
    子切片和切片究竟会不会影响，就抓住一点，它们是不是还共享数组？
    就是如果它们结构没有变化，那肯定是供共享的；
    但是结构变化了，就可能不是共享了
    比如从原来的切片中切出一个子切片，只需要看原本的切片有没有扩过容纳，后面的切片有没有扩过容，但凡有一个扩容了，它们就不会共享数组了
    所以不建议使用子切片，或者使用的时候是只读的，不会去改

    
## lib

go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql

