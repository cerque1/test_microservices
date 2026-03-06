package persistent

import (
	"context"
	"service-1/internal/entities"
	"service-1/internal/repo"
)

type FilmsUCImpl struct {
	filmsRepo repo.FilmsRepo
}

func NewFilmsUC(filmsRepo repo.FilmsRepo) *FilmsUCImpl {
	return &FilmsUCImpl{filmsRepo: filmsRepo}
}

func (fuc *FilmsUCImpl) Create(ctx context.Context, film entities.Film) (int, error) {
	return fuc.filmsRepo.Create(ctx, film)
}

func (fuc *FilmsUCImpl) GetById(ctx context.Context, id int) (entities.Film, error) {
	return fuc.filmsRepo.GetById(ctx, id)
}
