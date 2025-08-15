package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5, 6}
	soma(a)

}

func soma(si []int) {
	total := 0
	for _, v := range si {
		total = total + v
	}
	fmt.Println(total)
}
