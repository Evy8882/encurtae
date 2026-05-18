package handlers

import (
	"encoding/json"
	"encurtae/internal/repository"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
)

func putHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// logica pra editar a URL encurtada, usando o ID privado para identificar o documento no Firestore e atualizar a URL original associada a ele
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	var body struct {
		OriginalUrl string `json:"originalUrl"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if body.OriginalUrl == "" {
		http.Error(w, "Missing originalUrl", http.StatusBadRequest)
		return
	}

	firebaseService, err := repository.NewFirebaseService()
	if err != nil {
		http.Error(w, "Failed to initialize Firebase", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
    client, err := firebaseService.App.Firestore(ctx)
    if err != nil {
        http.Error(w, "Failed to initialize Firestore", http.StatusInternalServerError)
        return
    }
    defer client.Close()

	_, err = client.Collection("urls").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "originalUrl",
			Value: body.OriginalUrl,
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update URL: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "URL updated successfully")
}
