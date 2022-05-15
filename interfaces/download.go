package main

import (
	"fmt"
	"github.com/huahua/learngo/interfaces/testing"
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

type retriever interface {
	Get(url string) string
}

func GetRetriever() retriever {
	return testing.Retriever{}
}

func main() {
	r := GetRetriever()
	fmt.Println(r.Get("https://www.qq.com"))
}
