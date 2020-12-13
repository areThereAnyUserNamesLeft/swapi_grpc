package main

import (
	"context"
	"fmt"
	"grpc-swapi/generated/films"
	"grpc-swapi/infra/database/mongo"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	db, err := mongo.NewDBConnection(ctx, "mongodb://127.0.0.1:27017", "swapi")
	if err != nil {
		log.Fatal(err)
	}
	app := SwapiServiceServer{
		DB: db,
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	films.RegisterFilmsServiceServer(grpcServer, &app)
	fmt.Println("grpc listening on 8090")
	grpcServer.Serve(lis)
}

// SwapiServiceServer.
type SwapiServiceServer struct {
	DB   *mongo.Connection
	GRPC *films.UnimplementedFilmsServiceServer
}

// RequestFilm fufils interface from pb.
func (s *SwapiServiceServer) RequestFilm(ctx context.Context, in *films.FilmRequest) (*films.FilmsResponse, error) {
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

func (s *SwapiServiceServer) RequestFilms(ctx context.Context, in *films.FilmRequest) (*films.FilmsResponse, error) {
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

func (s *SwapiServiceServer) SearchFilms(ctx context.Context, in *films.SearchRequest) (*films.FilmsResponse, error) {
	return s.GRPC.SearchFilms(ctx, in)
}
