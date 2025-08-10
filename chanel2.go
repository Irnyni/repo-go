package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("goroutines:", runtime.NumGoroutine())
	canal := make(chan int)
	canal2 := make(chan int)
	wg.Add(2)
	go colocar(canal)
	fmt.Println("Goroutines :", runtime.NumGoroutine())
	go retirar(canal, canal2)
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	go func() {
		wg.Wait()
		close(canal2)
	}()
	for v := range canal2 {
		fmt.Printf("%v ", v)
	}
	fmt.Println("Goroutines :", runtime.NumGoroutine())
}

func colocar(canal chan int) {
	defer close(canal)
	defer wg.Done()
	for i := 0; i < 200; i++ {
		fmt.Println("Goroutines :", runtime.NumGoroutine())
		canal <- i
	}

}
func retirar(canal, canal2 chan int) {
	defer wg.Done()
	for v := range canal {
		canal2 <- v
	}

}
