package main

import "fmt"

func main() {
	dados := map[string][]string{
		"malengo": []string{
			"surf", "bola,", "caminhada",
		},
		"joão": []string{
			"nadar", "passear",
		},
		"carlos": []string{
			"dormir", "comer",
		},
	}

	for i, v := range dados {
		fmt.Println("\n", i)
		for _, item := range v {

			println(item)

		}

	}
	dados["debora"] = []string{

		"passear",
		"cozinhar",
	}
	for i, v := range dados {
		fmt.Println("\n", i)
		for _, item := range v {

			println(item)

		}

	}

	delete(dados, "debora")
	for i, v := range dados {
		fmt.Println("\n", i)
		for _, item := range v {

			println(item)

		}

	}
}
