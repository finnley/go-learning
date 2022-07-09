package main

import "fmt"

type User struct {
	Name string
	age  int
}

// 结构体接收器
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.age = newAge
}

func main() {
	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		age:  10,
	}
	u.ChangeName("Tom Changed!")
	u.ChangeAge(100)
	// {Tom 100}
	fmt.Printf("%v\n", u)

	// 因为 up 是 指针，所以内部的数据是可以改变的
	up := &User{
		Name: "Jerry",
		age:  12,
	}
	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!")
	up.ChangeAge(120)
	// &{Jerry 120}
	fmt.Printf("%v\n", up)
}

/**
指针接收器方法会影响到自身的状态，结构体不会
*/
