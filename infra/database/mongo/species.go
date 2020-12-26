package mongo

import (
	"context"
	"fmt"
	"grpc-swapi/generated/species"
	"grpc-swapi/structs"

	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
)

// GetSpeci finds a speci by index and populates a var it is passed.
func (c *Connection) GetSpeci(ctx context.Context, pk interface{}, f *species.Speci) error {

	cur, err := c.DB.Collection("species").Find(ctx, bson.M{"pk": pk})
	if err != nil {
		return err
	}

	sf := structs.Species{}
	if err = cur.All(ctx, &sf); err != nil {
		return err
	}
	fmt.Println(sf)

	return hydrateSpeci(0, sf, f)
}

// GetSpecies finds a speci by index and populates a var it is passed.
func (c *Connection) GetSpecies(ctx context.Context, pk []int32) ([]*species.Speci, error) {
	ff := []*species.Speci{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("species").Find(ctx, d)
	if err != nil {
		return ff, err
	}

	sf := structs.Species{}
	if err = cur.All(ctx, &sf); err != nil {
		return ff, err
	}

	for i := range sf {
		if contains(sf[i].Pk, pk) {
			var f species.Speci
			ff = append(ff, &f)
			err = hydrateSpeci(i, sf, &f)
			if err != nil {
				return ff, err
			}
		}
	}

	return ff, err
}

// SearchSpecies finds a speci by index and populates a var it is passed.
func (c *Connection) SearchSpecies(ctx context.Context, searchText string) ([]*species.Speci, error) {
	ff := []*species.Speci{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("species").Find(ctx, d)
	if err != nil {
		return ff, err
	}

	sf := structs.Species{}
	if err = cur.All(ctx, &sf); err != nil {
		return ff, err
	}

	for i := range sf {
		if matches(sf[i].Fields.Name, searchText) {
			var f species.Speci
			ff = append(ff, &f)
			err = hydrateSpeci(i, sf, &f)
			if err != nil {
				return ff, err
			}
		}
	}

	return ff, err
}

func hydrateSpeci(index int, ss structs.Species, s *species.Speci) error {
	edited, err := ptypes.TimestampProto(ss[index].Fields.Edited)
	if err != nil {
		return err
	}
	s.Edited = edited
	created, err := ptypes.TimestampProto(ss[index].Fields.Created)
	if err != nil {
		return err
	}
	s.Created = created
	s.Persons = itoI32(ss[index].Fields.People)
	s.Clasification = ss[index].Fields.Classification
	s.Name = ss[index].Fields.Name
	s.Designation = ss[index].Fields.Designation
	s.EyeColors = ss[index].Fields.EyeColors
	s.SkinColors = ss[index].Fields.SkinColors
	s.Language = ss[index].Fields.Language
	s.HairColors = ss[index].Fields.HairColors
	s.AverageLifespan = ss[index].Fields.AverageLifespan
	s.AverageHeight = ss[index].Fields.AverageHeight

	s.PK = int32(ss[index].Pk)
	return err
}
