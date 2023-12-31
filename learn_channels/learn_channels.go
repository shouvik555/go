package main

import "fmt"

func main() {
	/*in := make(chan string)
	out := make(chan string)
	go func() {
		name := <-in
		out <- fmt.Sprint("Hello ," + name)
	}()
	in <- "Bob"
	close(in)
	message := <-out
	fmt.Println(message)*/
	out := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(localI int) {
			out <- localI * 2
		}(i)
	}
	var result []int
	for i := 0; i < 10; i++ {
		val := <-out
		result = append(result, val)
	}
	fmt.Println(result)
}
