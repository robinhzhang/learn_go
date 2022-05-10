package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
func main() {
	var s []int //zero of value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)

	}
	fmt.Println(s)

	//copy
	s1 := []int{1, 24, 2, 3, 4, 2, 34, 4, 3, 2}
	s2 := make([]int, 10, 32)
	copy(s2, s1)
	fmt.Println(s2)

	//delete
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)

	//pop front
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println(s2)

	//pop back
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	fmt.Println(s2)

}
