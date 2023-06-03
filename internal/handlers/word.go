package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/mvahedii/wordbox/internal/services"
	"net/http"
)

type WordHandler struct {
	wordService services.WordService
	validator   *validator.Validate
}

func NewWordHandler(w services.WordService, v *validator.Validate) *WordHandler {
	return &WordHandler{
		wordService: w,
		validator:   v,
	}
}

func (repo *WordHandler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
}

type WordCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Def   string `json:"def" validate:"required"`
}
