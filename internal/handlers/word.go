package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/mvahedii/wordbox/internal/handlers/utils"
	"github.com/mvahedii/wordbox/internal/services"
	"github.com/mvahedii/wordbox/internal/services/dto"
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
	ctx := r.Context()
	var word WordCreateRequest

	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		utils.RespondBadRequest(w, err)
		return
	}

	wDTO := toWordDTO(word)

	err = repo.wordService.Create(ctx, wDTO)
	if err != nil {
		utils.RespondInternalServerError(w, err)
		return
	}

	//TODO: change this data
	utils.RespondOK(w, "test")
}

type WordCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Def   string `json:"def" validate:"required"`
}

func toWordDTO(w WordCreateRequest) *dto.Word {
	return &dto.Word{
		Title: w.Title,
		Def:   w.Def,
	}
}
