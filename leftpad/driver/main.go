package main

import (
	"fmt"
	"leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Internatiolization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😀'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '😀'))
	fmt.Println(leftpad.FormatRune("Internatiolization", 15, '😀'))
}
