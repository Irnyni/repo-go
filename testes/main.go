package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := soma(a)
	fmt.Println(b)
}

func soma(si []int) int {
	total := 0
	for _, v := range si {
		total = total + v
	}

	return total
}
func Soma2(a ...int) int {
	total := 0
	for _, v := range a {
		total = total + v
	}

	return total
}
