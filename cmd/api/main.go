package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Servidor Go rodando com sucesso!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
