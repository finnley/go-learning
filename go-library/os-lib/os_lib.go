package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func main() {
	//envs := os.Environ()
	//fmt.Println(envs)

	fmt.Println(runtime.GOOS)
	fmt.Println(GetExecDir()) // /Users/finnley/go/src/go-learning
}

func GetExecDir() string {
	if "darwin" == runtime.GOOS {
		a, _ := filepath.Abs(".")
		return a
	}
	p, err := os.Readlink("/proc/self/exe")
	if nil != err {
		panic("getRootDir got err:" + err.Error())
	}
	return path.Dir(p)
}
