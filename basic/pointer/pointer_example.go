package main

import "fmt"

//值传递 没用
//func swap(a, b  int ){
//	b,a =a,b
//}
//改,swap更好的实现
func swap_gai(a, b int) (int, int) {
	return b, a
}

// 指针传递
func swap1(a, b *int) {
	*b, *a = *a, *b
}

// 指针类型
func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	//指针不能运算

	//参数传递
	//go语言是值传递
	//穿值类型和指针类型需要根据情况来看
	a1, b1 := 3, 4
	swap1(&a1, &b1)
	fmt.Println(a1, b1)
	fmt.Println(*&a)
	a1, b1 = swap_gai(a1, b1)
	fmt.Println(a1, b1)
}
