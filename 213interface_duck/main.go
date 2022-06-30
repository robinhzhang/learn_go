package main

import (
	"fmt"
	"github.com/huahua/learngo/213interface_duck/mock"
	"github.com/huahua/learngo/213interface_duck/real"
	"time"
)

type Retriever interface {
	GetURLContent(url string) string
}

//215 Poster
type Poster interface {
	Post(url string, form map[string]string) string
}

//接口组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.GetURLContent("https://www.qq.com")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name": "huahua",
	})
}

func session(s RetrieverPoster) string {

	s.Post("http://www.imooc.com", map[string]string{
		"content": "faked immoc.com",
	})
	return s.GetURLContent("https://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "fake content"}
	inspect(r)
	fmt.Println(download(r))
	r = &real.Retriever{
		UserAgent: "Chrome",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//fmt.Println(download(r))
	// 214查看interface类型type assertion
	RealRetriever := r.(*real.Retriever)
	fmt.Println(RealRetriever.TimeOut)
	if MockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(MockRetriever.Content)
	} else {
		fmt.Println("not a mock retriever")
	}
	var mockRetriever mock.Retriever
	fmt.Println("try a session")
	fmt.Println(session(&mockRetriever))

}

// 214查看interface类型switch
func inspect(r Retriever) {
	fmt.Println("inspecting...")
	fmt.Printf(" > %v %T", r, r)
	fmt.Println(" > Type switch")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println(v.Content)
	case *real.Retriever:
		fmt.Println(v.UserAgent)
	}

}
