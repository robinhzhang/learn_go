package main

import (
	"fmt"
)

func lens(s string) int {
	//
	start := 0
	maxlength := 0
	lastOcc := make(map[rune]int)
	for i, ch := range []rune(s) {
		lastI, ok := lastOcc[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOcc[ch] = i
		fmt.Println(i, start, lastOcc)
	}
	return maxlength
}

func main() {
	fmt.Println(lens("abcabcbb"))
	fmt.Println(lens("Yes我爱笑着为我心碎的女孩"))
}
