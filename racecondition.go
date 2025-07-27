package main

import (
	"fmt"
	"sync"
)

var gr sync.WaitGroup
var i = 0
var mutex sync.Mutex

func main() {

	for i < 1000 {
		gr.Add(3)
		go cont1()
		go cont2()
		go cont3()

		gr.Wait()
	}

}

func cont1() {
	// mutex.Lock()
	i = i + 1
	// mutex.Unlock()
	fmt.Println(i)
	gr.Done()
}
func cont2() {
	// mutex.Lock()
	i = i + 1
	// mutex.Unlock()
	fmt.Println(i)
	gr.Done()
}
func cont3() {
	// mutex.Lock()
	i = i + 1
	// mutex.Unlock()
	fmt.Println(i)
	gr.Done()
}
