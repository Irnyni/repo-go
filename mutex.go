package main

import (
	"fmt"
	"runtime"
	"sync"
)

var goroutine sync.WaitGroup
var cont = 0
var mutex sync.Mutex

func main() {
	gerarGo(1000000)
	fmt.Println("Total do contador:", cont)

}

func gerarGo(num int) {
	goroutine.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			runtime.Gosched()
			mutex.Lock()
			cont += 1
			mutex.Unlock()
			goroutine.Done()
		}()

	}
	goroutine.Wait()
}
