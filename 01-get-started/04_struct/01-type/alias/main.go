package main

import "fmt"

type News struct {
	Name string
}

func (d News) Report() {
	fmt.Println("I am news: " + d.Name)
}

// 定义别名
type fakeNews = News

func main() {
	var n News = fakeNews{
		Name: "Hello",
	}
	n.Report()
}
