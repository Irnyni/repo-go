package main

import(
		"fmt"
		"encoding/json"

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

fs,err:= json.Marshal(funcionarios)
if err != nil{
	fmt.Println(err)
}
fmt.Println(string(fs))



}