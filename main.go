package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Listagem de Notas")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Exibindo uma nota...")
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Rejeitar a requisicao
		w.WriteHeader(405)
		return
	}
	fmt.Fprint(w, "Criando uma nota...")

}

func main() {
	fmt.Println("Servidor rodando na porta 5000")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
