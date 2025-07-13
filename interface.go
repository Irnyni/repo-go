package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}
type cliente struct {
	pessoa
	tipo   string
	saldos float64
}
type inadiplente interface {
	saldo(s float64)
}

func (c cliente) saldo(s float64) {
	fmt.Println(s)

}

func inadiplencia(i inadiplente) {
	if i.(cliente).saldos < 0 {
		fmt.Printf("Cliente inadiplente!!! Seu saldo devedor é de %v\n", i.(cliente).saldos)
	} else {
		fmt.Println(i.(cliente).saldos)
	}

}

func main() {
	cliente1 := cliente{pessoa{"julio", "cesar", 32}, "prime", -1500}
	fmt.Println(cliente1)
	cliente1.saldo(3450)
	inadiplencia(cliente1)
	fmt.Println(cliente1.saldos)

}
