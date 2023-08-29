package main

import "fmt"

func setTo10(a *int){
	*a = 10
}

func main() {
	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	//var z *int
	// fmt.Println(z) // => <nil>
	//fmt.Println(*z) // => panic: runtime error: invalid memory address or nil pointer dereference

	z := new(int)
	fmt.Println(z)
	fmt.Println(*z)

	x := 20
	fmt.Println(x)
	setTo10(&x)
	fmt.Println(x)
}
