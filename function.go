package main

import "fmt"

func main() {

	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8))
	fmt.Println(soma2([]int{1, 2, 3, 4, 56, 7}))

}

func soma(n ...int) int {
	soma := 0
	for _, v := range n {
		soma += v

	}
	return soma
}

func soma2(slice []int) int {
	soma := 0
	for _, v := range slice {
		soma += v
	}
	return soma

}
