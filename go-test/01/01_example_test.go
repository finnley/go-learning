package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	//t.SkipNow()
	res := Print1to20()
	fmt.Println("hey")
	testPrint(t)
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func testPrint(t *testing.T) {
	fmt.Println("test")
}

//func TestPrint999(t *testing.T) {
//	t.Run("a1", func(t *testing.T) {
//		fmt.Println("a1")
//	})
//	t.Run("a2", func(t *testing.T) {
//		fmt.Println("a2")
//	})
//	t.Run("a3", func(t *testing.T) {
//		fmt.Println("a3")
//	})
//}

func TestPrint2(t *testing.T) {
	res := Print1to20()
	res++
	if res != 211 {
		t.Errorf("Test Print2 failed")
	}
}

func TestAll(t *testing.T) {
	t.Run("TestPrint", TestPrint)
	t.Run("TestPrint2", TestPrint2)
}

// 使用 TestMain 作为初始化test，并且使用 m.Run() 来调用其他 tests 可以完成一些需要初始化操作的 testing，比如数据库连接，文件打开，REST服务登录等
// 如果没有在 TestMain 中调用 m.Run() ,则除了 TestMain 以外的其他 tests 都不会被执行
func TestMain(m *testing.M) {
	fmt.Println("test main first")
	m.Run()
}
