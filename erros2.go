package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Comportamento []string
}

func main() {

	pessoa1 := Pessoa{"julio", "César", []string{"calculista", "esperto", "controlador"}}

	file, err := converttojson(pessoa1)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(file))
	a := Pessoa{}
	err1 := json.Unmarshal(file, &a)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(a)
}
func converttojson(p Pessoa) ([]byte, error) {

	file, err := json.Marshal(p)
	return file, err

}
