package main

import (
	"fmt"
	"utf8"
)

func main() {
	str := "dsjkdshdjsdh....js"
	fmt.Printf("String %s\nLenght: %d, Runes: %d\n", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)))
}
