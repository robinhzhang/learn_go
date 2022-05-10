package main

import "fmt"

func main() {
	//数组声明方式
	var a [4]float32 // 等价于：var arr2 = [4]float32{}
	fmt.Println(a)   // [0 0 0 0]
	var b = [5]string{"ruby", "王二狗", "rose"}
	fmt.Println(b)                          // [ruby 王二狗 rose  ]
	var c = [5]int{'A', 'B', 'C', 'D', 'E'} // byte
	fmt.Println(c)                          // [65 66 67 68 69]
	d := [...]int{1, 2, 3, 4, 5}            // 根据元素的个数，设置数组的大小
	fmt.Println(d)                          //[1 2 3 4 5]
	e := [5]int{4: 100}                     // [0 0 0 0 100]
	fmt.Println(e)
	f := [...]int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
	fmt.Println(f)
	g := []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(g)

	//访问元素
	fmt.Println(g[2])

	//数组长度
	fmt.Println(len(g))

	//range 遍历数组
	h := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range h { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)

	// _ 忽略元素
	for _, v := range a { //ignores index
		fmt.Println(v)
	}

	//多维数组
	j := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /*  第三行索引为 2 */
	}
	fmt.Println(j)

	//数组是值类型 Go中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。
	k := [...]string{"USA", "China", "India", "Germany", "France"}
	l := k // a copy of a is assigned to b
	l[0] = "Singapore"
	fmt.Println("a is ", k)
	fmt.Println("b is ", l)
	//数组的大小是类型的一部分。因此[5]int和[25]int是不同的类型。因此，数组不能被调整大小。不要担心这个限制，因为切片的存在是为了解决这个问题。
	//m := [3]int{5, 78, 8}
	//var n [5]int
	//m = n //not possible since [3]int and [5]int are distinct types

}
