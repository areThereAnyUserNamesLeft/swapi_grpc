package structs

import "time"

type Films []struct {
	Fields Film   `json:"fields" bson:"fields"`
	Model  string `json:"model" bson:"model"`
	Pk     int    `json:"pk" bson:"pk"`
}

type Film struct {
	Starships    []int     `json:"starships" bson:"starships"`
	Edited       time.Time `json:"edited" bson:"edited"`
	Vehicles     []int     `json:"vehicles" bson:"vehicles"`
	Planets      []int     `json:"planets" bson:"planets"`
	Producer     string    `json:"producer" bson:"producer"`
	Title        string    `json:"title" bson:"title"`
	Created      time.Time `json:"created" bson:"created"`
	EpisodeID    int       `json:"episode_id" bson:"episode_id"`
	Director     string    `json:"director" bson:"director"`
	ReleaseDate  string    `json:"release_date" bson:"release_date"`
	OpeningCrawl string    `json:"opening_crawl" bson:"opening_crawl"`
	Characters   []int     `json:"characters" bson:"characters"`
	Species      []int     `json:"species" bson:"species"`
}
