package main

import(
		"fmt"
		"encoding/json"

)
type Funcionario []struct {
	Nome    string `json:"Nome"`
	Cargo   string `json:"Cargo"`
	Salario int    `json:"Salario"`
}

func main(){
a:=[]byte(`[{"Nome":"carlo","Cargo":"torneiro","Salario":3500},{"Nome":"flavio","Cargo":"ajustador","Salario":3000},{"Nome":"frigo","Cargo":"encarregado","Salario":10000}]`)
var b Funcionario
err:=json.Unmarshal(a,&b)
if err != nil{

		fmt.Println(err)

}
fmt.Println(string(a))
fmt.Println(b)

}