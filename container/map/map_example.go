package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "huahua",
	}
	m1 := map[string]map[string]int{} //复合map

	fmt.Println(m)
	fmt.Println(m1)

	m2 := make(map[string]int)
	fmt.Println(m2)

	var m3 map[string]int
	fmt.Println(m3)
	//遍历 不保证顺序
	for k, v := range m {
		fmt.Println(k, v)
	}

	//map hashmap 无序
	hua, ok := m["name"]
	fmt.Println(hua, ok)
	//	取不存在的值
	hua1, ok := m["nn"]
	fmt.Println(hua1, ok)
	//判断用ok

	//删除
	delete(m, "name")
	hua, ok = m["name"]
	fmt.Println(hua, ok)

	//map使用哈希表 必须可以比较想等
	//除了slice map func 的内建类型的值都可以做key
	//struct不包含上述字段，也可以做key

}
