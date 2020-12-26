package grpc

import (
	"context"
	"grpc-swapi/generated/species"
)

// RequestSpecies fufils interface from pb.
func (s *SwapiServer) RequestSpecies(ctx context.Context, in *species.SpeciesRequest) (*species.SpeciesResponse, error) {
	pk := in.GetSpecies()
	ff, err := s.DB.GetSpecies(ctx, pk)
	if err != nil {
		return &species.SpeciesResponse{
			Species: ff,
		}, err
	}
	return &species.SpeciesResponse{
		Species: ff,
	}, nil
}

// SearchSpecies fufils interface from pb.
func (s *SwapiServer) SearchSpecies(ctx context.Context, in *species.SearchRequest) (*species.SpeciesResponse, error) {
	sr := in.GetSearchText()
	ff, err := s.DB.SearchSpecies(ctx, sr)
	if err != nil {
		return &species.SpeciesResponse{
			Species: ff,
		}, err
	}
	return &species.SpeciesResponse{
		Species: ff,
	}, nil
}
