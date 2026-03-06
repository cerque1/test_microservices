package repo

import (
	"context"
	"service-1/internal/entities"
)

type FilmsRepo interface {
	GetById(ctx context.Context, id int) (entities.Film, error)
	Create(ctx context.Context, film entities.Film) (int, error)
}
