package service

import (
	"context"
	"errors"
	"github.com/mvahedii/wordbox/internal/repositories"
	"github.com/mvahedii/wordbox/internal/repositories/models"
	"github.com/mvahedii/wordbox/internal/services/dto"
)

type WordService struct {
	wr repositories.WordRepository
}

func toDto(w *models.Word) *dto.Word {
	return &dto.Word{
		ID:        w.ID,
		Title:     w.Title,
		Def:       w.Def,
		CreatedAt: w.CreatedAt,
	}
}

func fromDto(w *dto.Word) *models.Word {
	return &models.Word{
		ID:        w.ID,
		Title:     w.Title,
		Def:       w.Def,
		CreatedAt: w.CreatedAt,
	}
}

func NewWordService(r repositories.WordRepository) *WordService {
	return &WordService{
		wr: r,
	}
}

func (repo *WordService) Create(ctx context.Context, wordDTO *dto.Word) error {
	w := fromDto(wordDTO)
	err := repo.wr.Create(ctx, w)
	if err != nil {
		return errors.New("error on repo create ")
	}
	return nil

}
