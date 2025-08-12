package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	texto := strings.NewReader("texto do novo file")
	io.Copy(file, texto)
}
