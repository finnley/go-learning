package mock

// 接口实现
type Retriever struct {
	Content string
}

func (r *Retriever) Get(url string) string {
	return r.Content
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Content = form["contents"]
	return "ok"
}
