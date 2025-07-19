package main

import (
	"encoding/json"
	"fmt"
)

type Carro struct {
	Marca       string
	Modelo      string
	Ano         string
	Fipe        float64
	Combustivel string
	Km          int
}

func main() {
	carro1 := Carro{
		Marca:       "fiat",
		Modelo:      "argo",
		Ano:         "2015",
		Fipe:        45000.70,
		Combustivel: "gasolina",
		Km:          120000,
	}
	carro2 := Carro{"volkswagen", "passat", "2020", 200000, "gasolina",
		150000}
	f, err := json.Marshal(carro1)
	if err != nil {
		fmt.Println(err)
	}
	p, err := json.Marshal(carro2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(f), string(p))
}
