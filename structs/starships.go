package structs

type Starships []struct {
	Fields Starship `json:"fields" bson:"fields"`
	Model  string   `json:"model" bson:"model"`
	Pk     int      `json:"pk" bson:"pk"`
}
type Starship struct {
	Pilots           []int  `json:"pilots" bson:"pilots"`
	MGLT             string `json:"MGLT" bson:"MGLT"`
	StarshipClass    string `json:"starship_class" bson:"starship_class"`
	HyperdriveRating string `json:"hyperdrive_rating" bson:"hyperdrive_rating"`
}
