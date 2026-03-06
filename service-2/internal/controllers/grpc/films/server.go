package films

import (
	"context"
	"log"
	"service-2/internal/entities"
	"time"

	filmsv1 "github.com/cerque1/films_protos/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverApi struct {
	filmsv1.UnimplementedService1Server
	films Films
}

type Films interface {
	Create(
		ctx context.Context,
		film entities.Film,
	) (newId int, err error)
	GetById(
		ctx context.Context,
		id int,
	) (film entities.Film, err error)
}

func Register(gRPCServer *grpc.Server, films Films) {
	filmsv1.RegisterService1Server(gRPCServer, &serverApi{films: films})
}

func (s *serverApi) Create(ctx context.Context, in *filmsv1.CreateFilmRequest) (*filmsv1.CreateFilmResponse, error) {
	log.Println("start create")
	if in.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad name")
	}

	if in.GetLength() < 0 {
		return nil, status.Error(codes.InvalidArgument, "Bad length")
	}

	releaseDate, err := time.Parse(time.DateOnly, in.GetReleaseDate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "parse release time error: %s", err)
	}

	newId, err := s.films.Create(ctx, entities.Film{Name: in.GetName(), Length: int(in.GetLength()), ReleaseDate: releaseDate})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create error: %s", err)
	}
	return &filmsv1.CreateFilmResponse{NewId: int32(newId)}, nil
}

func (s *serverApi) GetById(ctx context.Context, in *filmsv1.GetByIdRequest) (*filmsv1.GetByIdResponse, error) {
	log.Println("start get film")
	film, err := s.films.GetById(ctx, int(in.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get film error: %s", err)
	}
	return &filmsv1.GetByIdResponse{
		Id:          int32(film.Id),
		Name:        film.Name,
		Length:      int32(film.Length),
		ReleaseDate: film.ReleaseDate.Format(time.DateOnly),
	}, nil
}
