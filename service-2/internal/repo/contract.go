package repo

import (
	"context"
	"service-2/internal/entities"
)

type FilmsRepo interface {
	Create(ctx context.Context, film entities.Film) (int, error)
	GetById(ctx context.Context, id int) (entities.Film, error)
}

type FilmsRepoFactory interface {
	CreateNew() FilmsRepo
}
