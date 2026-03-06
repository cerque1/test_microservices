package entities

import "time"

type Film struct {
	Id          int
	Name        string
	Length      int
	ReleaseDate time.Time
}
