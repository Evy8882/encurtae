package handlers

import (
	"context"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type UrlDoc struct {
	OriginalUrl string `firestore:"originalUrl"`
	ShortUrl    string `firestore:"shortUrl"`
}

func getHandler(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		shortCode := strings.TrimPrefix(r.URL.Path, "/r/")
		if shortCode == "" {
			http.Error(w, "Missing short URL", http.StatusBadRequest)
			return
		}

		iter := client.Collection("urls").Where("shortUrl", "==", shortCode).Limit(1).Documents(ctx)
		defer iter.Stop()

		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				http.Error(w, "URL not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Error fetching URL", http.StatusInternalServerError)
			return
		}

		var urlDoc UrlDoc
		if err := doc.DataTo(&urlDoc); err != nil {
			http.Error(w, "Error parsing URL data", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, urlDoc.OriginalUrl, http.StatusMovedPermanently)
	}
}
