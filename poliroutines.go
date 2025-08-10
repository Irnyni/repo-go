package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	routine := 10
	c := make(chan int)
	wg.Add(routine)
	criago(routine, c)

	go func() {
		wg.Wait()
		close(c)

	}()
	for v := range c {
		fmt.Printf("\nNúmeros adicionados no canal:\t%v", v)
	}
}

func criago(n int, c chan int) {

	for i := 0; i < n; i++ {

		go func(n int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				c <- i * n

			}

		}(i)
	}
}
