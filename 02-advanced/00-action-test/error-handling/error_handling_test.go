package error_handling

import (
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

// 自己实现 error 接口
// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

func MyErrorString(t *testing.T) {

}
