package structs

type Vehicles []struct {
	Fields Vehicle `json:"fields" bson:"fields"`
	Model  string  `json:"model" bson:"model"`
	Pk     int     `json:"pk" bson:"pk"`
}
type Vehicle struct {
	VehicleClass string `json:"vehicle_class" bson:"vehicle_class"`
	Pilots       []int  `json:"pilots" bson:"pilots"`
}
