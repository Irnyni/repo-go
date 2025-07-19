package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	endereco  string
}

func mudeme(p *pessoa) {
	p.endereco = "novo endereço!!!"
}

func main() {
	pessoa1 := pessoa{"jão", "Paulo", "caçapava"}
	fmt.Println(pessoa1)
	mudeme(&pessoa1)
	fmt.Println(pessoa1)
}
