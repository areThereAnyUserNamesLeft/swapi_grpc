package mongo

import (
	"context"
	"fmt"
	"grpc-swapi/generated/films"
	"grpc-swapi/structs"

	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
)

// GetFilm finds a film by index and populates a var it is passed.
func (c *Connection) GetFilm(ctx context.Context, pk interface{}, f *films.Film) error {

	cur, err := c.DB.Collection("films").Find(ctx, bson.M{"pk": pk})
	if err != nil {
		return err
	}

	sf := structs.Films{}
	if err = cur.All(ctx, &sf); err != nil {
		return err
	}
	fmt.Println(sf)

	return hydrateFilm(0, sf, f)
}

// GetFilms finds a film by index and populates a var it is passed.
func (c *Connection) GetFilms(ctx context.Context, pk []int32) ([]*films.Film, error) {
	ff := []*films.Film{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("films").Find(ctx, d)
	if err != nil {
		return ff, err
	}

	sf := structs.Films{}
	if err = cur.All(ctx, &sf); err != nil {
		return ff, err
	}

	for i := range sf {
		if contains(sf[i].Pk, pk) {
			var f films.Film
			ff = append(ff, &f)
			err = hydrateFilm(i, sf, &f)
			if err != nil {
				return ff, err
			}
		}
	}

	return ff, err
}

// GetFilms finds a film by index and populates a var it is passed.
func (c *Connection) SearchFilms(ctx context.Context, searchText string) ([]*films.Film, error) {
	ff := []*films.Film{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("films").Find(ctx, d)
	if err != nil {
		return ff, err
	}

	sf := structs.Films{}
	if err = cur.All(ctx, &sf); err != nil {
		return ff, err
	}

	for i := range sf {
		if matches(sf[i].Fields.Title, searchText) {
			var f films.Film
			ff = append(ff, &f)
			err = hydrateFilm(i, sf, &f)
			if err != nil {
				return ff, err
			}
		}
	}

	return ff, err
}

func hydrateFilm(index int, sf structs.Films, f *films.Film) error {
	edited, err := ptypes.TimestampProto(sf[index].Fields.Edited)
	if err != nil {
		return err
	}
	f.Edited = edited
	created, err := ptypes.TimestampProto(sf[index].Fields.Created)
	if err != nil {
		return err
	}
	f.Created = created
	f.Characters = itoI32(sf[index].Fields.Characters)
	f.Director = sf[index].Fields.Director
	f.EpisodeID = int32(sf[index].Fields.EpisodeID)
	f.OpeningCrawl = sf[index].Fields.OpeningCrawl
	f.Planets = itoI32(sf[index].Fields.Planets)
	f.ReleaseDate = sf[index].Fields.ReleaseDate
	f.Species = itoI32(sf[index].Fields.Species)
	f.Starships = itoI32(sf[index].Fields.Starships)

	f.Title = sf[index].Fields.Title
	f.PK = int32(sf[index].Pk)
	return err
}
