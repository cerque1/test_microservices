package httputil

type FilmCreateRequest struct {
	Name        string `json:"name"`
	Length      int    `json:"length"`
	ReleaseDate string `json:"release_date"`
}

type FilmCreateResponse struct {
	NewId int `json:"new_id"`
}

type GetFilmByIdResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Length      int    `json:"length"`
	ReleaseDate string `json:"release_date"`
}
