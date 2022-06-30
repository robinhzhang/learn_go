package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("system_interfaces/a.txt")
	if err != nil {
		panic(err)
	}
	printFileContents(file)
	s := `sds
		ds
		dsd

		sd
		d
		d`
	printFileContents(strings.NewReader(s))

	fmt.Println(string(1))
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
