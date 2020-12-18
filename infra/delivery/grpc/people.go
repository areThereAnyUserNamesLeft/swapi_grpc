package grpc

import (
	"context"
	"grpc-swapi/generated/people"
)

// RequestFilm fufils interface from pb.
func (s *SwapiServer) RequestPerson(ctx context.Context, in *people.PeopleRequest) (*people.PeopleResponse, error) {
	pk := int(in.GetPK()[0])
	p := people.Person{}
	err := s.DB.GetPerson(ctx, pk, &p)
	pp := []*people.Person{}
	if err != nil {
		return &people.PeopleResponse{
			People: pp,
		}, err
	}
	pp = append(pp, &p)
	return &people.PeopleResponse{
		People: pp,
	}, nil
}

func (s *SwapiServer) RequestPeople(ctx context.Context, in *people.PeopleRequest) (*people.PeopleResponse, error) {
	pk := in.GetPK()
	pp, err := s.DB.GetPeople(ctx, pk)
	if err != nil {
		return &people.PeopleResponse{
			People: pp,
		}, err
	}
	return &people.PeopleResponse{
		People: pp,
	}, nil
}

func (s *SwapiServer) SearchPersons(ctx context.Context, in *people.SearchRequest) (*people.PeopleResponse, error) {
	sr := in.GetSearchText()
	pp, err := s.DB.SearchPersons(ctx, sr)
	if err != nil {
		return &people.PeopleResponse{
			People: pp,
		}, err
	}
	return &people.PeopleResponse{
		People: pp,
	}, nil
}
