package handlers

import (
	"context"
	"encurtae/internal/repository"
	"fmt"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// liberar qualquer origem
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// métodos permitidos
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// headers permitidos
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// responder preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Opa, o que voce ta fazendo aqui? bb")
}

func Init() {
	firebaseService, err := repository.NewFirebaseService()
	if err != nil {
		fmt.Println("Erro ao inicializar Firebase:", err)
		return
	}

	client, err := firebaseService.App.Firestore(context.Background())
	if err != nil {
		fmt.Println("Erro ao inicializar Firestore:", err)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/post", postHandler)
	mux.HandleFunc("/r/", getHandler(client))

	fmt.Println("Servidor rodando na porta 8080")

	err = http.ListenAndServe(":8080", enableCORS(mux))
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
