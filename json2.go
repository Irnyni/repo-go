package main

import (
		"fmt"
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
sliceofbytes:=[]byte(`{"Marca":"fiat","Modelo":"argo","Ano":"2015","Fipe":45000.7,"Combustivel":"gasolina","Km":130000}`)
var carros Carro
err := json.Unmarshal(sliceofbytes,&carros)
if err!= nil{
		fmt.Println(err)
}
fmt.Println(carros.Modelo)

}