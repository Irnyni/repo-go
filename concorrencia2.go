package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	funcaogoroutines(100)
	wg.Wait()
}
func funcaogoroutines(s int) {
	wg.Add(s)
	for i := s; i > 0; i-- {
		fmt.Printf("go routine: %v\n", i)
		wg.Done()
	}

}
