package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Pessoa struct {
		Nome          string
		Sobrenome     string
		comportamento []string
	}

	pessoa1 := Pessoa{"julio", "César", []string{"calculista", "esperto", "controlador"}}

	file, err := json.Marshal(pessoa1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	fmt.Println(string(file))

}
