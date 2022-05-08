package main

import (
	"fmt"
	"reflect"
)

func main() {
	//s := [2:6] [:5] [:] [2:]
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := arr[:5]
	fmt.Println(s1, s2)
	//切片是引用类型 修改slice的值会跟着修改本体的值
	fmt.Println("更改本体的值")
	fmt.Println(arr)
	fmt.Println(s1)
	s1[0] = 123
	fmt.Println("更改后。。。。。")
	fmt.Println(arr)
	fmt.Println(s1)
	//reslice 在slice上再切
	s3 := s1[1:3]
	fmt.Println(s1)
	fmt.Println(s3)
	//slice扩展 可以向后扩展不能向前扩展 s[i] i<len(s) 扩展<cap(s)
	s4 := s1[3:6]
	fmt.Println(arr)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s4, len(s4), cap(s4))
	//空切片
	fmt.Println("空切片")
	var nilslice []int
	fmt.Println(nilslice == nil)
	fmt.Println(nilslice)
	fmt.Println(reflect.TypeOf(nilslice))
	// 注意两者区别
	s5 := []int{1, 3, 4}
	fmt.Println(s5)
	fmt.Println(reflect.TypeOf(s5))
	fmt.Println(len(s5))
	fmt.Println(cap(s5))
	s6 := [...]int{1, 3, 4}
	fmt.Println(s6)
	fmt.Println(reflect.TypeOf(s6))
	var s7 = make([]int, 5, 7)
	fmt.Println(s7)
	fmt.Println(reflect.TypeOf(s7))
	fmt.Println(len(s7))
	fmt.Println(cap(s7))

	// append append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	// append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
	// 但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
	// 返回的slice数组指针将指向这个空间，而原 数组的内容将保持不变；其它引用此数组的slice则不受影响

}
