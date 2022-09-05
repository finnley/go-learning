package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {
	name string
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

func (r Retriever) RefreshData(f func(s string)) {
	f("你好")
}
