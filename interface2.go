package main

import (
	"fmt"
)

type humano interface {
	falar()
}

func dizeralgumacoisa(h humano) {
	fmt.Println("esta falando")
}

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	p1 := Pessoa{"Julio", 21}
	p1.falar()
	dizeralgumacoisa(&p1)
}

func (p *Pessoa) falar() {
	fmt.Println(p.Nome, "diz oi!")
}
