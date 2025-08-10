package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	canal := make(chan int)
	canal2 := make(chan int)
	wg.Add(2)
	go colocar(canal)
	go retirar(canal, canal2)
	go func() {
		wg.Wait()
		close(canal2)
	}()
	for v := range canal2 {
		fmt.Println(v)
	}

}

func colocar(canal chan int) {
	defer close(canal)
	defer wg.Done()
	for i := 0; i < 200; i++ {
		canal <- i
	}

}
func retirar(canal, canal2 chan int) {
	defer wg.Done()
	for v := range canal {
		canal2 <- v
	}

}
