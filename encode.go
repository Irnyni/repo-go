package main

import (
	"os"
	"encoding/json"
)

type Carro struct {
	Marca       string  `json:"Marca"`
	Modelo      string  `json:"Modelo"`
	Ano         string  `json:"Ano"`
	Fipe        float64 `json:"Fipe"`
	Combustivel string  `json:"Combustivel"`
	Km          int     `json:"Km"`
}

func main(){
carro1 := Carro{
		Marca:       "fiat",
		Modelo:      "argo",
		Ano:         "2015",
		Fipe:        45000.70,
		Combustivel: "gasolina",
		Km:          130000,
	}
	carro2 := Carro{"volkswagen", "passat", "2020", 200000, "gasolina",
		150000}
encoder:=json.NewEncoder(os.Stdout)
encoder.Encode(carro1)
encoder.Encode(carro2)


}