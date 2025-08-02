package main

import (
	"fmt"
)

func main() {

	canal := make(chan int)
	go adicionar(canal, 10)
	for v := range canal {
		fmt.Println(v)
	}

}

func adicionar(c chan<- int, i int) {

	for t := i; t > 0; t-- {
		c <- t

	}
	close(c)

}
