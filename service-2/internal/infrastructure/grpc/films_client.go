package grpc

import (
	"context"
	"fmt"
	"service-2/internal/entities"
	"service-2/internal/repo"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	filmsv1 "github.com/cerque1/films_protos/gen"
)

type FilmsClient struct {
	api filmsv1.Service1Client
}

func NewFilmsClient(addr string) (*FilmsClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("Error connect films client: %s", err)
	}

	grpcClient := filmsv1.NewService1Client(conn)
	return &FilmsClient{api: grpcClient}, nil
}

func (fcli *FilmsClient) Create(ctx context.Context, film entities.Film) (int, error) {
	resp, err := fcli.api.Create(ctx, &filmsv1.CreateFilmRequest{Name: film.Name, Length: int32(film.Length), ReleaseDate: film.ReleaseDate.Format(time.DateOnly)})
	if err != nil {
		return -1, fmt.Errorf("Error create film: %s", err)
	}

	return int(resp.GetNewId()), nil
}

func (fcli *FilmsClient) GetById(ctx context.Context, id int) (entities.Film, error) {
	resp, err := fcli.api.GetById(ctx, &filmsv1.GetByIdRequest{Id: int32(id)})
	if err != nil {
		return entities.Film{}, fmt.Errorf("Error get film: %s", err)
	}

	releaseDate, err := time.Parse(time.DateOnly, resp.GetReleaseDate())
	if err != nil {
		return entities.Film{}, fmt.Errorf("Bad release date format: %s", err)
	}

	return entities.Film{
		Id:          int(resp.GetId()),
		Name:        resp.GetName(),
		Length:      int(resp.GetLength()),
		ReleaseDate: releaseDate,
	}, nil
}

type FilmsClientsFactoryImpl struct {
	addr string
}

func NewFilmsClientsFactory(addr string) *FilmsClientsFactoryImpl {
	return &FilmsClientsFactoryImpl{addr: addr}
}

func (fcf *FilmsClientsFactoryImpl) CreateNew() repo.FilmsRepo {
	// костыль но для примера пойдет
	filmsClient, _ := NewFilmsClient(fcf.addr)
	return filmsClient
}
