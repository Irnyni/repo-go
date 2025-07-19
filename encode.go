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
encoder:=json.NewEncoder(os.Stdout)
encoder.Encode(carro1)


}