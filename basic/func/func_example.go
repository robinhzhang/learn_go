package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//多返回值
func f(a, b int, op string) (int, error) {
	if op == "err" {
		return 0, fmt.Errorf("err")
	}

	return a + b, nil
}

func f1(a, b int) (c, d int) {
	fmt.Println(a, b)
	return 1, 2
}

//函数式编程，函数里面可以有函数
func apply1(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数
func sum1(numbers ...int) int {
	s := 0
	for i := 0; i < len(numbers); i++ {
		s += numbers[i]
	}
	return s
}
func main() {
	q, _ := f1(1, 2)
	fmt.Println(q)
	//error处理
	num, err := f(1, 2, "o")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	fmt.Println(apply1(pow, 22, 2))
	// 匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 2, 4))

	fmt.Println(sum1(1, 2, 3, 44, 23, 2))
}
