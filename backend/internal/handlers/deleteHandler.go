package handlers

import (
	"fmt"
	"net/http"

	"encurtae/internal/repository"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// conecta ao firebase
	firebaseService, err := repository.NewFirebaseService()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Firebase: %v", err), http.StatusInternalServerError)
		return
	}

	// pega o id que vai deletar
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	// deletar do Firestore usando o id
	ctx := r.Context()
	err = firebaseService.DeleteUrl(ctx, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete URL: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "URL deleted successfully")
}
