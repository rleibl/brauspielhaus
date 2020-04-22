package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

type Yeast struct {
	Name   string `json:name`
	Type   string `json:type`
	Amount string `json:amount` // "1 pkg" or something
}

type Malt struct {
	Name   string `json:name`
	EBC    string `json:ebc`
	Amount int    `json:amount` // amount in g
}

type Hop struct {
	Name   string `json:name`
	Alpha  string `json:alpha`
	Amount int    `json:amount` // amount in g
	Time   int    `json:time`   // time to cook this hops
}

type IngredientsOther struct {
	Name   string `json:name`
	Amount int    `json:amount` // amount in g
	Text   string `json:text`
}

type MashStep struct {
	Temperature int `json:temperature`
	Duration    int `json:duration`
}

type FermentationMeasurement struct {
	Date time.Time `json:date`
	Brix float32   `json:brix`
}

type Beer struct {
	// Housekeeping
	Id          int       `json:id`
	Name        string    `json:name`
	Description string    `json:description`
	Brewdate    time.Time `json:brewdate`

	// Mashing / Boiling
	Hops        []Hop              `json:hops`
	Malts       []Malt             `json:malts`
	IngredOther []IngredientsOther `json:ingredother`
	Mash        []MashStep         `json:mash`
	Boil        int                `json:boil` // boiltime in minutes
	MashNotes   string             `json:mashnotes`

	// Fermenting
	Yeasts []Yeast `json:yeasts`
	// Original wort: Measurements[0]
	// Final wort:    Measurements[len(Measurements)-1]
	Measurements      []FermentationMeasurement `json:measurements`
	FermentationNotes string                    `json:fermentationnotes`

	// Bottling / Kegging
	Volume float32 `json:volume`

	ABV float32 `json:abv` // calculated

	// Notes
	Notes string `json:notes`
}

func (b *Beer) UpdateCalculatedFields() {
}

func LoadBeersFromJson(directory string) []Beer {

	files, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		panic(err)
	}

	beers := make([]Beer, len(files))

	for i, f := range files {
		beers[i] = LoadBeerFromJson(f)
	}

	return beers
}

func LoadBeerFromJson(filename string) Beer {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var beer Beer

	err = json.Unmarshal(content, &beer)

	return beer
}

func (b *Beer) Store(filename string) {
}

func (b *Beer) Dump() {
}

func (b *Beer) ToJson() string {
	by, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(by)
}

func PrintBeerExample() {
	b := Beer{
		Id:          1,
		Name:        "Example Brew",
		Description: "This is an example brew to show the data structure",
		Brewdate:    time.Now(),

		Hops: []Hop{
			Hop{
				Name:   "Hallertauer Perle",
				Amount: 21,
				Time:   20,
			},
		},
		Malts: []Malt{
			Malt{
				Name:   "Pilsner Malz",
				EBC:    "10",
				Amount: 5000,
			},
		},
		Mash: []MashStep{
			MashStep{
				Temperature: 30,
				Duration:    0,
			},
			MashStep{
				Temperature: 60,
				Duration:    50,
			},
			MashStep{
				Temperature: 78,
				Duration:    0,
			},
		},
		Boil: 90,

		Yeasts: []Yeast{
			Yeast{
				Name:   "Saflager S-34",
				Amount: "1 pkg",
			},
		},

		Measurements: []FermentationMeasurement{},

		Volume: 23,
	}

	by, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(by))
}
