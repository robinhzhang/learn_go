package main

import (
	"fmt"
	"github.com/huahua/learngo/oop/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
