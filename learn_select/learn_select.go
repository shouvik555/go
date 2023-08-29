package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*in := make(chan int, 1)
	out := make(chan int, 1)

	out <- 1

	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	}*/

	/*
		in := make(chan int)
		out := make(chan int)

		select {
		case in <- 2:
			fmt.Println("wrote 2 to in")
		case v := <-out:
			fmt.Println("read", v, "from out")
		default:
			fmt.Println("nothing else works")
		}*/

	workers := 10000
	pool := make(chan func(int) int, workers)
	for i := 0; i < workers; i++ {
		pool <- func(in int) int {
			time.Sleep(1 * time.Second)
			result := 2 * in
			return result
		}
	}

	var wg sync.WaitGroup
	totalStart := time.Now()
	count := 0
	for i := 0; i < 100000; i++ {
		start := time.Now()
		select {
		case f := <-pool:
			fmt.Println("processing", i)
			count++
			wg.Add(1)
			go func(in int) {
				out := f(in)
				fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
				pool <- f
				wg.Done()
			}(i)
		default:
			fmt.Println("rejecting request", i, "too busy")
		}

	}
	wg.Wait()
	fmt.Println("total processed:", count)
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
