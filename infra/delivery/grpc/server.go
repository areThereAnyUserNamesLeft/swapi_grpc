package grpc

import "grpc-swapi/infra/database/mongo"

// SwapiServer.
type SwapiServer struct {
	DB *mongo.Connection
}
