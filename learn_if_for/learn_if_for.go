package main

import "fmt"

func main() {
	a := 10
	if a > 5 {
		b := a/2
		fmt.Println("a is bigger than 5:", a, b)
	} else {
		fmt.Println("a is less than or equal to 5")
	}
	//if a = 0 { fmt.Println("yes")}  ==> non-boolean condition in if statement
	//if a = 5 { fmt.Println("yes")}  ==> syntax error: cannot use assignment a = 5 as value
	if b := a/2; b > 5 {
		fmt.Println("a is bigger than 5:", a, b)
	} else {
		fmt.Println("a is less than or equal to 5")
	}

	for i := 0; i < 10; i++ {
		if i > 7 {
			break
			//break & continue is same as python
		}
		fmt.Println(i)
	}

	i := 0
	for {
		fmt.Println(i)
		i = i + 1
		if i > 10 { break }
	}

	//s := "Hello, world!"
	s := "😀 🍎"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
