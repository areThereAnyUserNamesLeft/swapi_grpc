package database

import (
	"context"
	"grpc-swapi/generated/films"
)

type FilmConn interface {
	GetFilm(context.Context, interface{}, *films.Film) error
	GetFilms(context.Context, []interface{}) ([]*films.Film, error)
	GetSearch(context.Context, []interface{}) ([]*films.Film, error)
}
