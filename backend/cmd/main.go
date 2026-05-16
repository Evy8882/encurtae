package main

import (
	"encurtae/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	handlers.Init()

	fmt.Println("Servidor ta rodando")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
