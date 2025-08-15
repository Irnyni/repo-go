package main

import "fmt"

type ErrorCustom struct {
	mensagem string
}

func (m ErrorCustom) Error() string {
	return "Mensagem custom"
}

func main() {

	E := ErrorCustom{"OLA ESSE É O ERRO"}
	mensagem(E)

}

func mensagem(e error) {
	fmt.Printf("erro apresentado: %v\n", e.(ErrorCustom).mensagem)
	fmt.Printf("erro apresentado: %v\n", e)
}
