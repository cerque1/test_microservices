package usecases

import (
	"context"
	"service-2/internal/entities"
)

type FilmsUC interface {
	Create(ctx context.Context, film entities.Film) (int, error)
	GetById(ctx context.Context, id int) (entities.Film, error)
}
