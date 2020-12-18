package grpc

import (
	"context"
	"grpc-swapi/generated/films"
)

// RequestFilm fufils interface from pb.
func (s *SwapiServer) RequestFilm(ctx context.Context, in *films.FilmRequest) (*films.FilmsResponse, error) {
	pk := int(in.GetPK()[0])
	f := films.Film{}
	err := s.DB.GetFilm(ctx, pk, &f)
	ff := []*films.Film{}
	if err != nil {
		return &films.FilmsResponse{
			Films: ff,
		}, err
	}
	ff = append(ff, &f)
	return &films.FilmsResponse{
		Films: ff,
	}, nil
}

func (s *SwapiServer) RequestFilms(ctx context.Context, in *films.FilmRequest) (*films.FilmsResponse, error) {
	pk := in.GetPK()
	ff, err := s.DB.GetFilms(ctx, pk)
	if err != nil {
		return &films.FilmsResponse{
			Films: ff,
		}, err
	}
	return &films.FilmsResponse{
		Films: ff,
	}, nil
}

func (s *SwapiServer) SearchFilms(ctx context.Context, in *films.SearchRequest) (*films.FilmsResponse, error) {
	sr := in.GetSearchText()
	ff, err := s.DB.SearchFilms(ctx, sr)
	if err != nil {
		return &films.FilmsResponse{
			Films: ff,
		}, err
	}
	return &films.FilmsResponse{
		Films: ff,
	}, nil
}
