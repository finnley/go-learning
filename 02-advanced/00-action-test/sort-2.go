package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}

// 根目录下的文件按时间大小排序，从远到近
func SortFile(path, name string) (files ByModTime) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fis, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	files = make(ByModTime, len(fis)+10)
	j := 0
	for i, v := range fis {
		if strings.Contains(fis[i].Name(), name) {
			files[j] = v
			j++
		}
	}
	files = files[:j]

	sort.Sort(ByModTime(files))
	// for _, fi := range files {
	// 	fmt.Println(fi.Name())
	// }
	return
}

func sortFile() {
	files, _ := ioutil.ReadDir("./redis_slow_log")

	sort.Sort(ByModTime(files))
	for _, redisInstance := range files {
		fmt.Println(redisInstance.Name())
	}
}

// 返回当下时间的文件，并删除大于 5 个的文件，删除最早的，如果目录下没有文件，就自动创建
func DealWithFiles(path, name string) (filename string) {
	timestamp := time.Now().Format("20060102.150405")
	filename = path + name + "." + timestamp
	files := SortFile(path, name)
	// fmt.Println(path + files[len(files)-1].Name())
	if len(files) > 5 {
		for k, _ := range files[5:] {
			err := os.Remove(path + files[k].Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if len(files) == 0 {
		f, err := os.Create(filename)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Println(filename)
	return filename
}

func main() {
	//path := "/res/csv/"
	//name := "user.csv"

	//files := DealWithFiles(path, name)
	//fmt.Println(files)
	sortFile()
}
