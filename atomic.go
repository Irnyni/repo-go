package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var gr sync.WaitGroup
var i int32 = 0
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
	atomic.AddInt32(&i, 1)
	atomic.LoadInt32(&i)
	fmt.Println(i)
	gr.Done()
}
func cont2() {
	atomic.AddInt32(&i, 1)
	atomic.LoadInt32(&i)
	fmt.Println(i)
	gr.Done()
}
func cont3() {
	atomic.AddInt32(&i, 1)
	atomic.LoadInt32(&i)
	fmt.Println(i)
	gr.Done()
}
