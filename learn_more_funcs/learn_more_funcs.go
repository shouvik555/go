package main

import "fmt"

func addOne (a int) int {
	return a + 1
}

func addTwo (a int) int {
	return a + 2
}

func printOps(a int, f func(int) int) {
	fmt.Println(f(a))
}

func makeAdder(b int) func(int) int {
	return func (a int) int {
		return a + b
	}
}
//Above func is returning a func

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}
//Above func accepts a func and returns a func

func main() {
	myAddOne := addOne
	fmt.Println(addOne(1))
	fmt.Println(myAddOne(1))
	// cant declare func inside a func, if you want to then
	// below method (anonymous function) needs to be used
	myAddOneNew := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddOneNew(2))
    //Below we are passing a numeric value and the function to call
	//At printOps, it uses the passed function as f with the param a
	printOps(1, addOne)
	printOps(2, addTwo)

	addOneA := makeAdder(1)
	addTwoA := makeAdder(2)
	fmt.Println(addOneA(1))
	fmt.Println(addTwoA(1))

	doubleAddOne := makeDoubler(addOneA)
	fmt.Println(doubleAddOne(1))
}
