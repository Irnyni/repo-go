package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 2, 9, 1, 5, 6}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)
    sabores := [] string{
		"banana",
		"uva",
		"abacate",
		"morango",
		"manga",
		"melancia",
		"laranja",
		"pitanga",
	}
	sort.Strings(sabores)
	fmt.Println(sabores)

	
}