package main

import "fmt"

func variable() {
	// 三种自定义变量形式
	var a = "s"
	fmt.Println(a)
	b := "ss"
	fmt.Println(b)
	var c int = 2
	fmt.Println(c)
	// 多变量声明
	d, e, f := 1, 2, 3
	fmt.Println(d, e, f)
	//批量自定义
	var (
		h int = 1
		i int = 2
	)
	fmt.Println(h, i)

	//const 常量 程序运行时，不会被修改
	const g = 1
	fmt.Println(g)
	//常量组
	const (
		k = 8989
		l //注意l输出
	)
	fmt.Println(k, l)

	//iota 每用一次就加1
	const (
		m = iota
		n
		p
	)
	fmt.Println(m, n, p)

}
