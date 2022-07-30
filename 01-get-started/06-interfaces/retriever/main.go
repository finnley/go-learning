package main

import (
	"fmt"
	"go-learning/01-get-started/06-interfaces/retriever/mock"
	"go-learning/01-get-started/06-interfaces/retriever/real"
	"time"
)

// 接口定义
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// go 接口由使用者来定义，Retriever是个接口
func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

const url = "https://www.imooc.com"

func post(poster Poster) {
	poster.Post(url, map[string]string{"name": "tech", "course": "golang"})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := &mock.Retriever{"this is a fake imooc.com"}
	r = retriever

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	// err
	//mockRetriever := r.(mock.Retriever)
	//fmt.Println(mockRetriever.Content)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Content)
	}

	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

/**
总结：
接口变量里面有什么？
1.实现者的类型
2.实现者的指针（也可以叫实现者的值），指针指向实现者

使用时千万不要用接口变量的地址，因为接口本身就自带了一个指针
1.接口变量自带指针
2.接口变量同样采用值传递，因为接口变量肚子里个指针，指向了实现着，所以几乎不需要使用接口的指针
3.指针接受者实现智能以指针方式使用；值接受都可以
*/
