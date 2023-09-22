package models

type LocalFood struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}
