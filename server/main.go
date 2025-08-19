package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/testando", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ola sou do servidor")
	})
	http.HandleFunc("/pagina_inicial", paginaInicial)
	http.HandleFunc("/videos", videos)
	http.HandleFunc("/sinopse", sinopses)

	fmt.Println("server iniciado")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println(err)
	}

}
func paginaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PAGINA INICIAL")
}

func videos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PAGINA DE VIDEOS")
}

func sinopses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "SINOPSE")
}
