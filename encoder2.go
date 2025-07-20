package main

import(
		"fmt"
		"encoding/json"
		"os"

)
type Funcionario struct{

		Nome string
		Cargo string
		Salario float64
}

func main(){

	f1:= Funcionario{"carlo","torneiro",3500}
	f2:= Funcionario{"flavio","ajustador",3000}
	f3:= Funcionario{"frigo","encarregado",10000}

funcionarios:= []Funcionario{f1,f2,f3}
fmt.Println(funcionarios)
err:=json.NewEncoder(os.Stdout).Encode(funcionarios)
if err != nil{
	fmt.Println(err)
}

}