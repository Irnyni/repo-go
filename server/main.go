package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Video struct {
	Id        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Imagem    string `json:"imagem"`
}

func main() {
	var err error

	db, err = sql.Open("sqlite3", "./videos.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS videos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			titulo TEXT,
			descricao TEXT,
			imagem TEXT
	)`)
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/pagina_inicial", paginaInicial)
	http.HandleFunc("/videos", videos)
	http.HandleFunc("/sinopse", sinopses)
	fmt.Println("server iniciado")
	err = http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println(err)
	}

}
func paginaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PAGINA INICIAL")
}

func videos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var v Video
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			fmt.Println(err)
		}

		_, err = db.Exec(`INSERT INTO videos(Titulo,Descricao,Imagem) 
		VALUES(?,?,?)`, v.Titulo, v.Descricao, v.Imagem)
		if err != nil {
			fmt.Println("Erro ao salvar no banco")
		}
	case http.MethodGet:
		rows, err := db.Query(`SELECT id,titulo, descricao, imagem FROM videos`)
		if err != nil {
			fmt.Println(err)
			defer rows.Close()

		}
		var videos []Video
		for rows.Next() {
			var v Video
			err := rows.Scan(&v.Id, &v.Titulo, &v.Descricao, &v.Imagem)
			if err != nil {
				http.Error(w, "Erro ao ler dados do banco", http.StatusInternalServerError)
				return
			}
			videos = append(videos, v)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videos)

	}

}

func sinopses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "SINOPSE")
}
