package entities

import "time"

type Film struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Length      int       `json:"length"`
	ReleaseDate time.Time `json:"release_date"`
}
