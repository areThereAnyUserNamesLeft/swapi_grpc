package structs

type Vehicles []struct {
	Fields struct {
		VehicleClass string        `json:"vehicle_class" bson:"vehicle_class"`
		Pilots       []interface{} `json:"pilots" bson:"pilots"`
	} `json:"fields" bson:"fields"`
	Model string `json:"model" bson:"model"`
	Pk    int    `json:"pk" bson:"pk"`
}
