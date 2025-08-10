package main

import (
	"fmt"
)

func main() {

	c := gen()

	receive(c)

	fmt.Println("sair")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {

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

}
