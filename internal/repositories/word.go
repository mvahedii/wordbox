package repositories

import (
	"context"
	"github.com/mvahedii/wordbox/internal/repositories/models"
)

type WordRepository interface {
	Create(ctx context.Context, word models.Word)
}
