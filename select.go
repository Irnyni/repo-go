package main

import "fmt"

func main() {
	x := 10
	a := make(chan int, 10)
	b := make(chan int, 10)
	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i

		}

	}(x)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i

		}

	}(x)

	for i := 0; i < x*2; i++ {
		select {

		case v := <-a:
			fmt.Println("canal A: ", v)

		case v := <-b:
			fmt.Println("canal b: ", v)

		}
	}
}
