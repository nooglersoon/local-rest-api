package mocks

import "github.com/nooglersoon/locale-rest-api/pkg/models"

var LocalFoods []models.LocalFood = []models.LocalFood{
	{
		Id:          1,
		Name:        "Sate Jando",
		Address:     "Gedung Sate",
		Latitude:    107.62,
		Longitude:   -61.200,
		Description: "Sate with cattle fat",
	},
	{
		Id:          2,
		Name:        "Nasi Goreng Ato",
		Address:     "Taruna Bakti",
		Latitude:    107.62,
		Longitude:   -61.200,
		Description: "Nasi goremg kampung style",
	},
}
