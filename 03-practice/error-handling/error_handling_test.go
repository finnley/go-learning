package error_handling

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	// 当打开文件失败时会返回一个非 nil 值 的 error
	// 这里打开文件失败时会打印错误消息并终止运行
	_, err := os.Open("filename.ext")
	if err != nil {
		// 2022/08/23 09:54:00 open filename.ext: no such file or directory
		log.Fatal(err)
	}
	// do something with the open *File f
}

// 下面的代码在 errors/errors.go 中实现
// 自己实现 error 接口
// errorString is a trivial implementation of error.
//type errorString struct {
//	s string
//}
//
//func (e *errorString) Error() string {
//	return e.s
//}
//
//// New returns an error that formats as the given text.
//func New(text string) error {
//	return &errorString{text}
//}

/**
=== RUN   TestMyErrorString
math: square root of negative number
--- PASS: TestMyErrorString (0.00s)
PASS

*/
func TestMyErrorString(t *testing.T) {
	_, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}

	_, err = Sqrt2(-1)
	if err != nil {
		fmt.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// todo
	return 1, nil
}

func Sqrt2(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	// todo
	return 1, nil
}

func Sqrt3(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	// todo
	return 1, nil
}

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

func TestNegativeSqrt(t *testing.T) {

}

// demo:
func TestDemo3(t *testing.T) {
	err := os.Remove("/tmp/nonexist")
	log.Println(err)
}
