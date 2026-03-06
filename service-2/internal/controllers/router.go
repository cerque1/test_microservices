package controllers

import (
	"net/http"
	_ "service-2/docs"
	"service-2/internal/controllers/films"
	"service-2/internal/usecases"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title 			films list & test microservices
// @version			1.0
// @license.name 	Apache 2.0

// @host 			localhost:8080
// @BasePath 		/

// @accept 			json
// @produce 		json
func InitRouter(filmsUC usecases.FilmsUC) {
	filmsRouter := films.NewFilmsRouter(filmsUC)

	http.HandleFunc("GET /films/{id}", filmsRouter.GetById)
	http.HandleFunc("POST /films", filmsRouter.Create)

	http.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
