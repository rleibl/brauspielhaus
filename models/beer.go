package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
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
	Date string  `json:date`
	Brix float32 `json:brix`
}

type Beer struct {
	// Housekeeping
	Id          int    `json:id`
	Name        string `json:name`
	Description string `json:description`
	Brewdate    string `json:brewdate`
	Status      string `json:status` // Automatic

	// Mashing / Boiling
	Hops        []Hop              `json:hops`
	Malts       []Malt             `json:malts`
	IngredOther []IngredientsOther `json:ingredother`
	Mash        []MashStep         `json:mash`
	Boil        int                `json:boil` // boiltime in minutes
	MashNotes   string             `json:mashnotes`

	// Fermenting
	FirstFermStart string  `json:firstfermstart` // brewdate, unlesse set otherwise
	Yeasts         []Yeast `json:yeasts`

	// Original wort: Measurements[0]
	// Final wort:    Measurements[len(Measurements)-1]
	Measurements      []FermentationMeasurement `json:measurements`
	FermentationNotes string                    `json:fermentationnotes`

	// Bottling / Kegging
	SecondFermStart string  `json:secondfermstart` // start date of second vermentation
	Volume          float32 `json:volume`

	ConditioningStart string `json:conditioningstart` // start date of conditioning
	Ready             string `json:ready`             // time brew is considered ready

	ABV float32 `json:abv` // calculated

	// Durations
	//
	// First Fermentation: FirstFermStart -> SecondFermStart
	DurFirstFerm int `json:durfirstferm` // Automatic
	// Second Fermentation: SecondFermStart -> ConditioningStart
	DurSecondFerm int `json:dursecondferm` // Automatic
	// ConditioningStart: ConditioningStart -> Ready
	DurConditioning int `json:durconditioning` // Automatic
	// Total: Brewdate -> Ready
	DurTotal int `json:durtotal` // Automatic

	// Notes
	Notes string `json:notes`
}

func (b *Beer) Validate() {

	if b.FirstFermStart == "" {
		b.FirstFermStart = b.Brewdate
	}

	b.Status = "Planned"
	if b.Brewdate != "" {
		b.Status = "First Fermentation"
	}

	if b.SecondFermStart != "" {
		b.DurFirstFerm = Duration(b.FirstFermStart, b.SecondFermStart)
		b.Status = "Second Fermentation"
		if b.ConditioningStart != "" {
			b.DurSecondFerm = Duration(b.SecondFermStart, b.ConditioningStart)
			b.Status = "Conditioning"
			if b.Ready != "" {
				b.DurConditioning = Duration(b.ConditioningStart, b.Ready)
				b.Status = "Ready"
				b.DurTotal = Duration(b.Brewdate, b.Ready)
			}
		}
	}

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

	return FromJson(content)
}

func FromJson(b []byte) Beer {

	var beer Beer

	err := json.Unmarshal(b, &beer)
	if err != nil {
		panic(err)
	}

	return beer
}

func (b *Beer) ToJson() string {
	by, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(by)
}

func (b *Beer) Store(filename string) {
}

func (b *Beer) Print() {

	fmt.Printf("Id:     %d\n", b.Id)
	fmt.Printf("Name:   %s\n", b.Name)
	fmt.Printf("Desc:   %s\n", b.Description)
	fmt.Printf("Date:   %s\n", b.Brewdate)
	fmt.Printf("Status: %s\n", b.Status)

	fmt.Printf("Hops:\n")
	for i, h := range b.Hops {
		fmt.Printf("    %d %s (%d g, %d min)\n", i+1, h.Name, h.Amount, h.Time)
	}

	fmt.Printf("Malts:\n")
	for i, m := range b.Malts {
		fmt.Printf("    %d %s (%s EBC, %d Amount)\n", i+1, m.Name, m.EBC, m.Amount)
	}

	fmt.Printf("Mash Steps:\n")
	for i, m := range b.Mash {
		fmt.Printf("    %d:  %d Deg, %3d min\n", i+1, m.Temperature, m.Duration)
	}

	fmt.Printf("Boil Time: %d\n", b.Boil)

	fmt.Printf("Yeast:\n")
	for i, y := range b.Yeasts {
		fmt.Printf("    %d %s (%s)\n", i+1, y.Name, y.Amount)
	}

	fmt.Printf("Yield: %fl\n", b.Volume)

	fmt.Printf("First Fermentation:  %s (%d days)\n", b.FirstFermStart, b.DurFirstFerm)
	fmt.Printf("Second Fermentation: %s (%d days)\n", b.SecondFermStart, b.DurSecondFerm)
	fmt.Printf("Conditioning:        %s (%d days)\n", b.ConditioningStart, b.DurConditioning)
	fmt.Printf("Ready:               %s (%d days)\n", b.Ready, b.DurTotal)

}

func BeerTemplate() Beer {
	b := Beer{
		Id:          1,
		Name:        "Example Brew",
		Description: "This is an example brew to show the data structure",
		Brewdate:    "01.01.1970",

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

	return b
}

func PrintBeerTemplate() {
	b := BeerTemplate()
	fmt.Println(b.ToJson())
}
