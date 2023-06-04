package handlers

import (
	"github.com/gorilla/mux"
	"github.com/mvahedii/wordbox/internal/services"
	"net/http"
)

func New(wordService services.WordService) http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	//wordHandler := NewWordHandler

	//r.Methods(http.MethodGet).Path("/word").HandlerFunc()

	return r
}
