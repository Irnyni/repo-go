package main

import "fmt"

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go func(c chan int, q chan int, n int) {
		for n > 0 {
			fmt.Println("atual:", <-c)
			n--
		}
		quit <- 1
	}(canal, quit, 500)
	func(c chan int, q chan int) {
		a := 0
		for {

			select {
			case c <- a:
				a++
			case <-quit:
				return

			}

		}

	}(canal, quit)

}
