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
fmt.Println("Slice normal:",funcionarios)
fmt.Println("Saida os.Stdout: ")
err:=json.NewEncoder(os.Stdout).Encode(funcionarios)
if err != nil{
	fmt.Println(err)
}
fmt.Println("Saída após encoder não salvando o dado:\n", funcionarios)
a,err:= json.Marshal(funcionarios)
if err != nil{
	fmt.Println(err)
}

fmt.Println("Saida após marshal, criando a variavel a: \n",string(a))


var f []Funcionario
sliceb:= []byte(a)
err1:= json.Unmarshal(sliceb,&f)
if err != nil{fmt.Println(err1)}
fmt.Println("Tranformação de json para go através do unmarshal:",f)




}