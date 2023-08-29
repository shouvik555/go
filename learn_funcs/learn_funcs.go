package main

import "fmt"


func main () {
	a := addNumbers(2,3)
	fmt.Println(a)

	a = addNumbers(4,10)
	fmt.Println(a)

	div, rem := divAndRem(2,3)
	fmt.Println(div, rem)

	// ignore the second value if you dont intend to use
	div, _ = divAndRem(2,3)
	fmt.Println(div)

}

func addNumbers(a int, b int) int {
	return a + b
}

func divAndRem(a int, b int) (int, int) {
	return a / b, a % b
}

// remember that the values of variables are not referenced, they are assigned to new address,
// so changes to variable or array inside a function won't update the same on called function