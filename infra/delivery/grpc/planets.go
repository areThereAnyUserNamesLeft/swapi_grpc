package grpc

import (
	"context"
	"grpc-swapi/generated/planets"
)

// RequestPlanet fufils interface from pb.
func (s *SwapiServer) RequestPlanet(ctx context.Context, in *planets.PlanetsRequest) (*planets.PlanetsResponse, error) {
	pk := int(in.GetPK()[0])
	f := planets.Planet{}
	err := s.DB.GetPlanet(ctx, pk, &f)
	ff := []*planets.Planet{}
	if err != nil {
		return &planets.PlanetsResponse{
			Planets: ff,
		}, err
	}
	ff = append(ff, &f)
	return &planets.PlanetsResponse{
		Planets: ff,
	}, nil
}

// RequestPlanets fufils interface from pb.
func (s *SwapiServer) RequestPlanets(ctx context.Context, in *planets.PlanetsRequest) (*planets.PlanetsResponse, error) {
	pk := in.GetPK()
	ff, err := s.DB.GetPlanets(ctx, pk)
	if err != nil {
		return &planets.PlanetsResponse{
			Planets: ff,
		}, err
	}
	return &planets.PlanetsResponse{
		Planets: ff,
	}, nil
}

// SearchPlanets fufils interface from pb.
func (s *SwapiServer) SearchPlanets(ctx context.Context, in *planets.SearchRequest) (*planets.PlanetsResponse, error) {
	sr := in.GetSearchText()
	ff, err := s.DB.SearchPlanets(ctx, sr)
	if err != nil {
		return &planets.PlanetsResponse{
			Planets: ff,
		}, err
	}
	return &planets.PlanetsResponse{
		Planets: ff,
	}, nil
}
