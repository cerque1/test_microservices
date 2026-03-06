package usecases

import (
	"context"
	"service-1/internal/entities"
)

type FilmsUC interface {
	GetById(ctx context.Context, id int) (entities.Film, error)
	Create(ctx context.Context, film entities.Film) (int, error)
}
