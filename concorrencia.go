package main

import (
	"fmt"
)

func main() {
	segunda()
	primeira()

}

func primeira() {
	for i := 0; i < 200; i++ {
		fmt.Println(i)
	}

}

func segunda() {
	for a := 200; a > 0; a-- {
		fmt.Println(a)

	}

}
