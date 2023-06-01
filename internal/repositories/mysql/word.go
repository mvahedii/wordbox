package mysql

import (
	"context"
	"github.com/mvahedii/wordbox/internal/repositories"
	"github.com/mvahedii/wordbox/internal/repositories/models"
)

type Word struct {
	ID    uint
	Title string
	Def   string
}

type WordRepo struct {
}

func NewWordRepo(ctx context.Context) repositories.WordRepository {
	return &WordRepo{}
}

func (w WordRepo) Create(ctx context.Context, word models.Word) {

}
