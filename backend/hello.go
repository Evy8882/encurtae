package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Opa, o que voce ta fazendo aqui?")
}

func main() {

	http.HandleFunc("/", homeHandler)

	fmt.Println("Servidor ta rodando")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
