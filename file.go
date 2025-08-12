package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Println("erro:", err)
		return
	}
	texto := strings.NewReader("texto do novo file")
	io.Copy(file, texto)
}
