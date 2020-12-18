package mongo

import (
	"context"
	"fmt"
	"grpc-swapi/generated/people"
	"grpc-swapi/structs"

	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
)

// GetFilm finds a film by index and populates a var it is passed.
func (c *Connection) GetPerson(ctx context.Context, pk interface{}, p *people.Person) error {

	cur, err := c.DB.Collection("people").Find(ctx, bson.M{"pk": pk})
	if err != nil {
		return err
	}

	sp := structs.People{}
	if err = cur.All(ctx, &sp); err != nil {
		return err
	}
	fmt.Println(sp)

	return hydratePerson(0, sp, p)
}

// GetFilms finds a film by index and populates a var it is passed.
func (c *Connection) GetPeople(ctx context.Context, pk []int32) ([]*people.Person, error) {
	pp := []*people.Person{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("people").Find(ctx, d)
	if err != nil {
		return pp, err
	}

	sp := structs.People{}
	if err = cur.All(ctx, &sp); err != nil {
		return pp, err
	}

	for i := range sp {
		if contains(sp[i].Pk, pk) {
			var p people.Person
			pp = append(pp, &p)
			err = hydratePerson(i, sp, &p)
			if err != nil {
				return pp, err
			}
		}
	}

	return pp, err
}

// GetFilms finds a film by index and populates a var it is passed.
func (c *Connection) SearchPersons(ctx context.Context, searchText string) ([]*people.Person, error) {
	pp := []*people.Person{}

	d := bson.D{{}}

	cur, err := c.DB.Collection("people").Find(ctx, d)
	if err != nil {
		return pp, err
	}

	sp := structs.People{}
	if err = cur.All(ctx, &sp); err != nil {
		return pp, err
	}

	for i := range sp {
		if matches(sp[i].Fields.Name, searchText) {
			var p people.Person
			pp = append(pp, &p)
			err = hydratePerson(i, sp, &p)
			if err != nil {
				return pp, err
			}
		}
	}

	return pp, err
}

func hydratePerson(index int, sp structs.People, p *people.Person) error {
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
	p.Birthyear = sp[index].Fields.BirthYear
	p.EyeColor = sp[index].Fields.EyeColor
	p.HairColor = sp[index].Fields.HairColor
	p.SkinColor = sp[index].Fields.SkinColor
	p.Gender = sp[index].Fields.Gender
	p.Height = sp[index].Fields.Height
	p.Mass = sp[index].Fields.Mass
	p.Homeworld = int32(sp[index].Fields.Homeworld)

	p.Name = sp[index].Fields.Name
	p.PK = int32(sp[index].Pk)
	return err
}
