package main

import "fmt"

func main() {
	type Pessoa struct {
		nome      string
		sobrenome string
		sabores   []string
	}

	pessoa1 := Pessoa{
		nome:      "joao",
		sobrenome: "silva",
		sabores:   []string{"banana", "caju", "melão"},
	}

	pessoa2 := Pessoa{nome: "maria", sobrenome: "joana", sabores: []string{"morango", "chocolate", "abacaxi"}}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

}
