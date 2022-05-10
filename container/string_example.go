package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱笑着为我心碎的女孩"
	fmt.Println(len(s))
	//Unicode编码
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	//range遍历 pos会+3
	for i, ch := range s { // ch is rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	//获得字符数数量
	fmt.Println(utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	//这样操作就可以国际化字符串操作
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()

	//其他string操作 Filds空格

}
