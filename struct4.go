package main

import "fmt"

// Person struct with Name and Age fields

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {

	p1 := pessoa{"Robson", "Gomes", 24}
	p2 := pessoa{"Maria", "Joana", 30}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.apresentar()
	p2.apresentar()

}

func (p pessoa) apresentar() {
	fmt.Printf("Olá, meu nome é %s %s e tenho %d anos.\n", p.nome, p.sobrenome, p.idade)
}
