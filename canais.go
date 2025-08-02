package main

import (
	"fmt"
	"sync"
)

var gr sync.WaitGroup
var mutex sync.Mutex
var total = 0

func main() {
	canal := make(chan int, 3)
	a := 0

	for a < 1000 {
		gr.Add(3)

		go contar(canal)
		go contar(canal)
		go contar(canal)

		gr.Wait()

		for i := 0; i < 3; i++ {
			val := <-canal
			total += val
		}

		a++
	}

	fmt.Println("Total:", total)
}

func contar(canal chan int) {
	canal <- 1

	gr.Done()
}
