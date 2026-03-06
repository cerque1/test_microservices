package persistent

import (
	"context"
	"database/sql"
	"errors"
	"service-1/internal/entities"
	"time"
)

type FilmsRepoImpl struct {
	db *sql.DB
}

func NewFilmsRepo(db *sql.DB) *FilmsRepoImpl {
	return &FilmsRepoImpl{db: db}
}

func (fr *FilmsRepoImpl) Create(ctx context.Context, film entities.Film) (int, error) {
	row := fr.db.QueryRowContext(ctx, "INSERT INTO films(name, length, release_date) VALUES($1, $2, $3) RETURNING id", film.Name, film.Length, film.ReleaseDate)
	var newId int
	err := row.Scan(&newId)
	if err != nil {
		return -1, errors.New("insert film error")
	}
	return newId, nil
}

func (fr *FilmsRepoImpl) GetById(ctx context.Context, id int) (entities.Film, error) {
	row := fr.db.QueryRowContext(ctx, "SELECT * FROM films WHERE id = $1", id)
	film := entities.Film{}
	var releaseDateStr string
	err := row.Scan(&film.Id,
		&film.Name,
		&film.Length,
		&releaseDateStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Film{}, errors.New("no such record to select")
		}
		return entities.Film{}, err
	}
	var releaseDate time.Time
	if releaseDate, err = time.Parse(time.RFC3339, releaseDateStr); err != nil {
		return entities.Film{}, errors.New("error parse release date")
	}
	film.ReleaseDate = releaseDate
	return film, nil
}
