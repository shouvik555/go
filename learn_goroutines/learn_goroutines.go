package main

import (
	"fmt"
	"sync"
)

/*func runMe() {
	fmt.Println("Hello from goroutine")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		runMe()
		wg.Done()
	}()
	wg.Wait()
}*/

func runMe(name string) {
	fmt.Println("Hello to", name, "from a goroutine")
}

func main() {
	var wg sync.WaitGroup
	/*wg.Add(1)
	go func(name string) {
		runMe(name)
		wg.Done()
	}("Bob")
	wg.Wait()*/
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(localI int) {
			fmt.Println(localI)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
