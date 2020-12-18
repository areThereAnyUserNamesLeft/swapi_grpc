package main

import (
	"context"
	"fmt"
	"grpc-swapi/generated/films"
	"grpc-swapi/generated/people"
	"grpc-swapi/infra/database/mongo"
	delivery "grpc-swapi/infra/delivery/grpc"
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
	app := delivery.SwapiServer{
		DB: db,
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	films.RegisterFilmsServiceServer(grpcServer, &app)
	people.RegisterPeopleServiceServer(grpcServer, &app)
	fmt.Println("grpc listening on 8090")
	grpcServer.Serve(lis)
}
