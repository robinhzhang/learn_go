package main

import (
	"fmt"
)

func main() {
	//经典款 for循环
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	//for {} 与while效果相同

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	// for condition{}
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	//for range
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	//break continue和其他语言差不多
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}
