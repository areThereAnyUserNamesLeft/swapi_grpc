package structs

import "time"

type Planets []struct {
	Fields struct {
		Edited         time.Time `json:"edited" bson:"edited"`
		Climate        string    `json:"climate" bson:"climate"`
		SurfaceWater   string    `json:"surface_water" bson:"surface_water"`
		Name           string    `json:"name" bson:"name"`
		Diameter       string    `json:"diameter" bson:"diameter"`
		RotationPeriod string    `json:"rotation_period" bson:"rotation_period"`
		Created        time.Time `json:"created" bson:"created"`
		Terrain        string    `json:"terrain" bson:"terrain"`
		Gravity        string    `json:"gravity" bson:"gravity"`
		OrbitalPeriod  string    `json:"orbital_period" bson:"orbital_period"`
		Population     string    `json:"population" bson:"population"`
	} `json:"fields" bson:"fields"`
	Model string `json:"model" bson:"model"`
	Pk    int    `json:"pk" bson:"pk"`
}
