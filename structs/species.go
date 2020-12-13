package structs

import "time"

type Species []struct {
	Fields struct {
		Edited          time.Time `json:"edited" bson:"edited"`
		Classification  string    `json:"classification" bson:"classification"`
		Name            string    `json:"name" bson:"name"`
		Designation     string    `json:"designation" bson:"designation"`
		Created         time.Time `json:"created" bson:"created"`
		EyeColors       string    `json:"eye_colors" bson:"eye_colors"`
		People          []int     `json:"people" bson:"people"`
		SkinColors      string    `json:"skin_colors" bson:"skin_colors"`
		Language        string    `json:"language" bson:"language"`
		HairColors      string    `json:"hair_colors" bson:"hair_colors"`
		Homeworld       int       `json:"homeworld" bson:"homeworld"`
		AverageLifespan string    `json:"average_lifespan" bson:"average_lifespan"`
		AverageHeight   string    `json:"average_height" bson:"average_height"`
	} `json:"fields" bson:"fields"`
	Model string `json:"model" bson:"model"`
	Pk    int    `json:"pk" bson:"pk"`
}
