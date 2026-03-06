package persistent

import (
	"context"
	"log"
	"service-2/internal/entities"
	"service-2/internal/repo"
)

type FilmsUCImpl struct {
	filmsRepoFactory repo.FilmsRepoFactory
}

func NewFilmsUc(filmsRepoFactory repo.FilmsRepoFactory) *FilmsUCImpl {
	return &FilmsUCImpl{filmsRepoFactory: filmsRepoFactory}
}

func (fuc *FilmsUCImpl) Create(ctx context.Context, film entities.Film) (int, error) {
	log.Println("start create")
	client := fuc.filmsRepoFactory.CreateNew()
	return client.Create(ctx, film)
}

func (fuc *FilmsUCImpl) GetById(ctx context.Context, id int) (entities.Film, error) {
	log.Println("start get by id")
	client := fuc.filmsRepoFactory.CreateNew()
	return client.GetById(ctx, id)
}
