package structs

import "time"

type Transports []struct {
	Fields Transport `json:"fields" bson:"fields"`
	Model  string    `json:"model" bson:"model"`
}
type Transport struct {
	Edited               time.Time `json:"edited" bson:"edited"`
	Consumables          string    `json:"consumables" bson:"consumables"`
	Name                 string    `json:"name" bson:"name"`
	Created              time.Time `json:"created" bson:"created"`
	CargoCapacity        string    `json:"cargo_capacity" bson:"cargo_capacity"`
	Passengers           string    `json:"passengers" bson:"passengers"`
	MaxAtmospheringSpeed string    `json:"max_atmosphering_speed" bson:"max_atmosphering_speed"`
	Crew                 string    `json:"crew" bson:"crew"`
	Length               string    `json:"length" bson:"length"`
	Model                string    `json:"model" bson:"model"`
	CostInCredits        string    `json:"cost_in_credits" bson:"cost_in_credits"`
	Manufacturer         string    `json:"manufacturer" bson:"manufacturer"`
	Pk                   int       `json:"pk" bson:"pk"`
}
