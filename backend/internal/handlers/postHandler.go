package handlers

import (
	"encoding/json"
	"encurtae/internal/repository"
	"fmt"
	"net/http"

	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func postHandler(w http.ResponseWriter, r *http.Request) {
	firebaseService, err := repository.NewFirebaseService()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Firebase: %v", err), http.StatusInternalServerError)
		return
	}

	// pegar originalUrl: string, enviado na requisição POST
	var body struct {
		OriginalUrl string `json:"originalUrl"`
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	originalUrl := body.OriginalUrl

	// gerar shortURL, criando um ID único, usando a função generateShortURL() e atribuindo o resultado à variável shortUrl
	var shortUrl string
	shortUrl, err = generateShortURL()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate short URL: %v", err), http.StatusInternalServerError)
		return
	}

	if originalUrl == "" {
		http.Error(w, "Missing originalUrl", http.StatusBadRequest)
		return
	}

	if shortUrl == "" {
		http.Error(w, "Failed to generate short URL", http.StatusInternalServerError)
		return
	}

	// gerar ID único para alterações e exclusões (privado, não exposto na API)
	id, err := generateShortURL()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate ID: %v", err), http.StatusInternalServerError)
		return
	}
	// salvar no Firestore
	ctx := r.Context()
	client, err := firebaseService.App.Firestore(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Firestore: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = client.Collection("urls").Doc(id).Set(ctx, map[string]interface{}{
		"originalUrl": originalUrl,
		"shortUrl":    shortUrl,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save URL mapping: %v", err), http.StatusInternalServerError)
		return
	}

	//retornar objeto JSON com id, shortUrl, e originalUrl

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":          id,
		"shortUrl":    fmt.Sprintf("http://localhost:8080/r/%s", shortUrl), // Substituir pelo domínio real
		"originalUrl": originalUrl,
	})
}

func generateShortURL() (string, error) {
	code := make([]byte, 12)

	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		code[i] = charset[n.Int64()]
	}

	return string(code), nil
}
