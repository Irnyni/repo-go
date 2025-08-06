package main

import (
	"fmt"
	"sync"
)

func main() {

	canal1 := make(chan int)
	canal2 := make(chan int)
	canais := make(chan int)
	go colocar(canal1, canal2)
	go merge(canal1, canal2, canais)

	for v := range canais {
		fmt.Println(v)
	}

}

func merge(c1, c2, juntar chan int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(c1, juntar chan int) {

		defer wg.Done()
		for v := range c1 {

			juntar <- v

		}
	}(c1, juntar)
	go func(c2, juntar chan int) {

		defer wg.Done()
		for v := range c2 {

			juntar <- v

		}

	}(c2, juntar)
	wg.Wait()
	close(juntar)
}
func colocar(c1, c2 chan int) {
	a := 100
	for i := 0; i < a; i++ {
		if i%2 == 0 {
			c1 <- i
		} else {
			c2 <- i
		}

	}
	close(c1)
	close(c2)
}
