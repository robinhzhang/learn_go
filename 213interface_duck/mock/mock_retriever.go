package mock

import "fmt"

type Retriever struct {
	Content string
}

func (r Retriever) GetURLContent(url string) string {
	return r.Content
}

func (r *Retriever) Post(url string, form map[string]string) string {
	//注意指针接受者
	r.Content = form["content"]
	return "ok"
}

//类似于tostring（）方法
func (r Retriever) String() string {
	return fmt.Sprintf("Retriever:{content=%s}", r.Content)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	//注意指针接受者
	r.Content = form["content"]
	return "ok"
}

//类似于tostring（）方法
func (r Retriever) String() string {
	return fmt.Sprintf("Retriever:{content=%s}", r.Content)
}
