package main

import (
	"fmt"
	"sync"
)

var gr sync.WaitGroup

func main() {
	gr.Add(2)
	c := gen()

	go receive(c)
	gr.Wait()
	fmt.Println("sair")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		defer gr.Done()
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i
		}

	}()

	return c
}

func receive(channel <-chan int) {

	for v := range channel {
		fmt.Println(v)
	}
	defer gr.Done()
}
