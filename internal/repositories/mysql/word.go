package mysql

import (
	"context"
	"errors"
	"github.com/mvahedii/wordbox/internal/repositories"
	"github.com/mvahedii/wordbox/internal/repositories/models"
	"gorm.io/gorm"
	"time"
)

type Word struct {
	ID        uint
	Title     string
	Def       string
	CreatedAt time.Time
}

func fromModel(w *models.Word) *Word {
	return &Word{
		ID:        w.ID,
		Title:     w.Title,
		Def:       w.Def,
		CreatedAt: w.CreatedAt,
	}
}

func toModel(w *Word) *models.Word {
	return &models.Word{
		ID:        w.ID,
		Title:     w.Title,
		Def:       w.Def,
		CreatedAt: w.CreatedAt,
	}
}

type WordRepo struct {
	db *gorm.DB
}

func NewWordRepo(db *gorm.DB) repositories.WordRepository {
	return &WordRepo{
		db: db,
	}
}

func (repo *WordRepo) Create(ctx context.Context, w *models.Word) error {
	word := fromModel(w)

	err := repo.db.Create(word).Error
	if err != nil {
		return errors.New("error on create")
	}

	w = toModel(word)

	return nil
}
