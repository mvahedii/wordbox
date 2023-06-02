package services

import (
	"context"
	"github.com/mvahedii/wordbox/internal/services/dto"
)

type WordService interface {
	Create(ctx context.Context, w *dto.Word) error
}
