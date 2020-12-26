package mongo

import (
	"context"
	"fmt"
	"grpc-swapi/generated/planets"
	"grpc-swapi/structs"

	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
)

// GetPlanet finds a planet by index and populates a var it is passed.
func (c *Connection) GetPlanet(ctx context.Context, pk interface{}, p *planets.Planet) error {

	cur, err := c.DB.Collection("planets").Find(ctx, bson.M{"pk": pk})
	if err != nil {
		return err
	}

	sp := structs.Planets{}
	if err = cur.All(ctx, &sp); err != nil {
		return err
	}
	fmt.Println(sp)

	return hydratePlanet(0, sp, p)
}

// GetPlanets finds a planet by index and populates a var it is passed.
func (c *Connection) GetPlanets(ctx context.Context, pk []int32) ([]*planets.Planet, error) {
	pp := []*planets.Planet{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("planets").Find(ctx, d)
	if err != nil {
		return pp, err
	}

	sp := structs.Planets{}
	if err = cur.All(ctx, &sp); err != nil {
		return pp, err
	}

	for i := range sp {
		if contains(sp[i].Pk, pk) {
			var p planets.Planet
			pp = append(pp, &p)
			err = hydratePlanet(i, sp, &p)
			if err != nil {
				return pp, err
			}
		}
	}

	return pp, err
}

// GetPlanets finds a planet by index and populates a var it is passed.
func (c *Connection) SearchPlanets(ctx context.Context, searchText string) ([]*planets.Planet, error) {
	pp := []*planets.Planet{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("planets").Find(ctx, d)
	if err != nil {
		return pp, err
	}

	sp := structs.Planets{}
	if err = cur.All(ctx, &sp); err != nil {
		return pp, err
	}

	for i := range sp {
		if matches(sp[i].Fields.Name, searchText) {
			var p planets.Planet
			pp = append(pp, &p)
			err = hydratePlanet(i, sp, &p)
			if err != nil {
				return pp, err
			}
		}
	}

	return pp, err
}

func hydratePlanet(index int, sp structs.Planets, p *planets.Planet) error {
	edited, err := ptypes.TimestampProto(sp[index].Fields.Edited)
	if err != nil {
		return err
	}
	p.Edited = edited
	created, err := ptypes.TimestampProto(sp[index].Fields.Created)
	if err != nil {
		return err
	}
	p.Created = created
	p.Climate = sp[index].Fields.Climate
	p.SurfaceWater = sp[index].Fields.SurfaceWater
	p.Name = sp[index].Fields.Name
	p.Diameter = sp[index].Fields.Diameter
	p.RotationPeriod = sp[index].Fields.RotationPeriod
	p.Terrain = sp[index].Fields.Terrain
	p.Gravity = sp[index].Fields.Gravity
	p.OrbitalPeriod = sp[index].Fields.OrbitalPeriod
	p.Population = sp[index].Fields.Population

	p.Name = sp[index].Fields.Name
	p.PK = int32(sp[index].Pk)
	return err
}
