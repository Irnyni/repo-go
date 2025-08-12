package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	logs, err := os.Create("logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logs.Close()
	log.SetOutput(logs)

	file2, err := os.Open("naoexiste.txt")
	if err != nil {
		log.Println("erro:", err)
	}
	defer file2.Close()
}
