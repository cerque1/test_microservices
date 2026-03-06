package films

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"service-2/internal/entities"
	"service-2/internal/httputil"
	"service-2/internal/usecases"
)

func createInternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	body, _ := json.Marshal(httputil.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf("Internal server error: %s", err)})
	w.Write(body)
}

func createBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	body, _ := json.Marshal(httputil.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("Bad request error: %s", err)})
	w.Write(body)
}

type FilmsRouter struct {
	filmsUC usecases.FilmsUC
}

func NewFilmsRouter(filmsUC usecases.FilmsUC) *FilmsRouter {
	return &FilmsRouter{filmsUC: filmsUC}
}

// Create film godoc
// @Summary		Create film
// @Description	Create film
// @Tags 		films
// @Acess  		json
// @Produce 	json
// @Param 		subscription 	body	httputil.FilmCreateRequest	true	"create film body"
// @Success 	200 {object}	httputil.FilmCreateResponse "created film id"
// @Failure		400 {object}	httputil.ErrorResponse
// @Failure		500 {object}	httputil.ErrorResponse
// @Router		/films [post]
func (fr *FilmsRouter) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		createInternalServerErrorResponse(w, err)
		return
	}

	createFilmRequest := httputil.FilmCreateRequest{}
	err = json.Unmarshal(body, &createFilmRequest)
	if err != nil {
		createBadRequestResponse(w, err)
		return
	}

	if createFilmRequest.Name == "" {
		createBadRequestResponse(w, errors.New("Bad film name"))
		return
	}
	if createFilmRequest.Length < 0 {
		createBadRequestResponse(w, errors.New("Bad film length"))
		return
	}

	release_date, err := time.Parse(time.DateOnly, createFilmRequest.ReleaseDate)
	if err != nil {
		createBadRequestResponse(w, fmt.Errorf("Bad film release date: %s", err))
		return
	}

	newId, err := fr.filmsUC.Create(r.Context(), entities.Film{Name: createFilmRequest.Name, Length: createFilmRequest.Length, ReleaseDate: release_date})
	if err != nil {
		createInternalServerErrorResponse(w, fmt.Errorf("Create film error: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(httputil.FilmCreateResponse{NewId: newId})
	w.Write(resp)
}

// Get film by id godoc
// @Summary		Get film
// @Description	Get film by id
// @Tags 		films
// @Produce 	json
// @Param 		id 	path		int	true	"film id"
// @Success 	200 {object}	httputil.GetFilmByIdResponse
// @Failure		400 {object}	httputil.ErrorResponse
// @Failure		500 {object}	httputil.ErrorResponse
// @Router		/films/{id} [get]
func (fr *FilmsRouter) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		createBadRequestResponse(w, fmt.Errorf("Bad film id"))
		return
	}

	film, err := fr.filmsUC.GetById(r.Context(), id)
	if err != nil {
		createInternalServerErrorResponse(w, fmt.Errorf("Get film error: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(httputil.GetFilmByIdResponse{
		Id:          film.Id,
		Name:        film.Name,
		Length:      film.Length,
		ReleaseDate: film.ReleaseDate.Format(time.DateOnly),
	})
	w.Write(resp)
}
