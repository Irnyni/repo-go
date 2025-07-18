package main

import "fmt"

func teste(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func slice(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total

}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(teste(1, 2, 3, 4, 5))
	fmt.Println(teste(a...))
	fmt.Println(slice(a))
}
