package handlers

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Opa, o que voce ta fazendo aqui? bb")
}

func Init() {
	http.HandleFunc("/", homeHandler)
}
