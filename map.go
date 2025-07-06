// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	x := make([]string, 26, 26)
	x = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo",
		"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba",
		"Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
		"Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(x)
	fmt.Printf("tamanho= %v\ncapacidade= %v", len(x), cap(x))
	for i := 0; i < cap(x); i++ {
		fmt.Println(x[i])
	}

	x = append(x, "portugal")
	fmt.Printf("tamanho= %v\ncapacidade= %v", len(x), cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
		fmt.Printf("índice %v = %v\n", i, x[i])
	}

}
