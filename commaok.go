package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 45; i < 100; i++ {
			c <- i

		}
	}()

	for v := range c {
		fmt.Println(v)
	}

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}
