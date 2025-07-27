package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go segunda()
	go primeira()
	wg.Wait()
}

func primeira() {
	for i := 0; i < 200; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func segunda() {
	for a := 200; a > 0; a-- {
		fmt.Println(a)

	}
	wg.Done()

}
