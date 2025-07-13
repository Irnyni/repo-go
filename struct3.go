package main

import "fmt"

func main() {
	type Pessoa struct {
		nome      string
		sobrenome string
		sabores   []string
	}

	pessoa1 := Pessoa{"felipe", "malengo", []string{"banana", "maça"}}
	pessoa2 := Pessoa{
		nome:      "maria",
		sobrenome: "joana",
		sabores: []string{
			"uva",
			"pera",
		},
	}
	fmt.Println(pessoa1)
	for _, v := range pessoa1.sabores {

		fmt.Println(v)
	}
	fmt.Println(pessoa2)
	for _, v := range pessoa2.sabores {

		fmt.Println(v)
	}
	x := []Pessoa{pessoa1, pessoa2}
	fmt.Println(x)
	for _, v := range x {
		for _, v := range v.sabores {
			fmt.Println(v)

		}

	}
}
