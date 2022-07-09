package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 将 Info 映射成数据库中的一张表，但是它能表述的信息是有限的
type Info struct {
	Name   string `json:"name"`// 比如Name是映射成mysql中的char类型还是varchar类型还是text类型？表述并不明确
	Age    int    `json:"age,omitempty"` // 不序列化零值
	Gender string `json:"-"` // 不序列化 {"name":"Finnley"}
}

// 取数使用反射包
// 使用自己定义的 tag
type Text struct {
	Name   string `orm:"name, max_length=17, min_length=5"`
	Age    int    `orm:"age, min=18, max=80"`
	Gender string `orm:"gender, required"`
}


func main() {
	// 结构体标签
	/**
	结构体的字段除了名字和类型外，还可以有一个可选的标签 (tag)；
	它是一个附属于字段的字符串，可以是文档或其他的重要标记。
	比如在我们解析json或生成json文件时，常用到的 encoding/json包，它提供了一些默认的标签，例如 omitempty 标签，可以在序列化的时候忽略0值或者空值，而标签的作用不进行序列化，其效果和直接将结构体中的字段写成小写的效果一样
	*/

	info := Info{
		Name:   "Finnley",
		Gender: "男",
	}
	re, _ := json.Marshal(info)
	fmt.Println(string(re)) // {"Name":"Finnley","Gender":"男"}
	fmt.Println(re) // [123 34 78 97 109 101 34 58 34 70 105 110 110 108 101 121 34 44 34 71 101 110 100 101 114 34 58 34 231 148 183 34 125]

	// 通过反射包识别tag
	text := Text{
		Name:   "夜",
		Age:    24,
		Gender: "女",
	}
	// TypeOf 获取text对象的类型
	t := reflect.TypeOf(text)
	fmt.Println("Type: ", t.Name()) // Struct 类型是Text
	fmt.Println("Kind: ", t.Kind()) // Kind:  struct

	// NumField 获取有多少属性
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // 获取结构体的每一个字段
		tag := field.Tag.Get("orm")
		fmt.Printf("%d.%v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		// todo 提取使用字符串处理
	}
	/**
	1.Name (string), tag: 'name, max_length=17, min_length=5'
	2.Age (int), tag: 'age, min=18, max=80'
	3.Gender (string), tag: 'gender, required'
	 */
}
