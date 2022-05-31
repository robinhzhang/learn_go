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

func download(r Retriever) string {
	return r.GetURLContent("https://www.qq.com")
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

}

// 214查看interface类型switch
func inspect(r Retriever) {
	fmt.Printf("%v %T", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println(v.Content)
	case *real.Retriever:
		fmt.Println(v.UserAgent)
	}

}
