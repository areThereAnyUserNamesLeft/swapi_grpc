package structs

import "time"

type People []struct {
	Fields struct {
		Edited    time.Time `json:"edited" bson:"edited"`
		Name      string    `json:"name" bson:"name"`
		Created   time.Time `json:"created" bson:"created"`
		Gender    string    `json:"gender" bson:"gender"`
		SkinColor string    `json:"skin_color" bson:"skin_color"`
		HairColor string    `json:"hair_color" bson:"hair_color"`
		Height    string    `json:"height" bson:"height"`
		EyeColor  string    `json:"eye_color" bson:"eye_color"`
		Mass      string    `json:"mass" bson:"mass"`
		Homeworld int       `json:"homeworld" bson:"homeworld"`
		BirthYear string    `json:"birth_year" bson:"birth_year"`
	} `json:"fields" bson:"fields"`
	Model string `json:"model" bson:"model"`
	Pk    int    `json:"pk" bson:"pk"`
}
