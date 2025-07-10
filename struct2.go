package main

import "fmt"

func main() {
	type Pessoa struct {
		nome      string
		sobrenome string
		endereco  string
		idade     int
	}

	type Profissao struct {
		Pessoa
		cargo    string
		salario  float64
		admissao string
	}

	pessoa1 := Profissao{
		Pessoa: Pessoa{
			nome:      "robson",
			sobrenome: "gomes",
			endereco:  "rua trevo",
			idade:     24,
		},
		cargo:    "Analista",
		salario:  3500.00,
		admissao: "20/02/2023",
	}
	fmt.Println(pessoa1)
	fmt.Printf("Funcionário: %v %v\nCargo: %v\nAdmitido em :%v", pessoa1.nome, pessoa1.sobrenome, pessoa1.cargo, pessoa1.admissao)

}
