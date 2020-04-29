package db

import (
	"errors"
	"github.com/rleibl/brauspielhaus/config"
	"github.com/rleibl/brauspielhaus/models"
)

var beers []models.Beer

func Init() {

	c := config.GetConfig()

	beers = models.LoadBeersFromJson(c.JsonPath)
	for _, b := range beers {
		b.Validate()
	}
}

func GetBeers() []models.Beer {
	return beers
}

func GetBeer(id int) (*models.Beer, error) {

	for _, b := range beers {
		if b.Id == id {
			return &b, nil
		}
	}

	return nil, errors.New("No such id")
}
