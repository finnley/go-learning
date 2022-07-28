package main

import (
	"fmt"
	"go-learning/01-get-started/06-interfaces/infra"
	"io/ioutil"
	"net/http"
)

func retrieve(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

func main1() {
	resp, err := http.Get("https://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", bytes)
}

func main2() {
	fmt.Println("%s", retrieve("https://www.imooc.com"))
}

func main3() {
	retriever := infra.Retriever{}
	fmt.Println(retriever.Get("https://www.imooc.com"))
}

//func getRetriever() infra.Retriever {
//	return infra.Retriever{}
//}

//func getRetriever() testing.Retriever {
//	return testing.Retriever{}
//}
//
//func main4() {
//	//var retriever infra.Retriever = getRetriever()
//	var retriever testing.Retriever = getRetriever()
//	//retriever := getRetriever()
//	fmt.Println(retriever.Get("https://www.imooc.com"))
//}

// ?: Something that can Get
type retriever interface {
	Get(url string) string
}

func getRetriever() retriever {
	//return testing.Retriever{}
	return infra.Retriever{}
}

func main() {
	var retriever = getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}
