package main

import "fmt"

func main() {
	type Pessoa struct {
		nome      string
		sobrenome string
		sabores   []string
	}

	pessoa1 := Pessoa{"Julio", "Santos", []string{"baunilha", "morango"}}
	pessoa2 := Pessoa{"Ana", "Silva", []string{"chocolate", "limão"}}
	pessoa3 := Pessoa{"Carlos", "Oliveira", []string{"flocos"}}
	pessoa4 := Pessoa{"Mariana", "Costa", []string{"morango", "coco", "menta"}}
	pessoa5 := Pessoa{"Pedro", "Souza", []string{"chocolate"}}

	lista := map[string]Pessoa{

		pessoa1.sobrenome: pessoa1,
		pessoa1.sobrenome: pessoa2,
		pessoa2.sobrenome: pessoa3,
		pessoa3.sobrenome: pessoa4,
		pessoa4.sobrenome: pessoa5,
	}

	for _, valor := range lista {

		fmt.Println(valor)
		for _, v := range valor.sabores {

			fmt.Println(v)
		}
	}

}
