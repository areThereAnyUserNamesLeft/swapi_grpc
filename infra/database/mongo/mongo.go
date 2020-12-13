package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewDBConnection creates a Connection.
func NewDBConnection(ctx context.Context, addr, DBname string) (*Connection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		return &Connection{}, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return &Connection{}, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return &Connection{}, err
	}
	db := client.Database(DBname)
	return &Connection{DB: db, Client: client}, err
}

// Connection is a convienient struct for a mongoDB connection.
type Connection struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func itoI32(ii []int) []int32 {
	out := []int32{}
	for i := range ii {
		out = append(out, int32(i))
	}
	return out
}

func i32sToStr(ii []int32) string {
	out := "["
	for _, v := range ii {
		out = fmt.Sprintf("%s\"%d\", ", out, int(v))
	}

	out = fmt.Sprintf("%s]", out[0:len(out)-2])
	fmt.Println(out)
	return out
}

func contains(i int, arr []int32) bool {
	for _, a := range arr {
		if int(a) == i {
			return true
		}
	}
	return false
}
