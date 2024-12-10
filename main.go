package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Rota Default")
}

func olaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ola Handler")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")
	mux := http.NewServeMux()
	//rota coringa /
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ola", olaHandler)

	http.ListenAndServe(":5000", mux)
}
