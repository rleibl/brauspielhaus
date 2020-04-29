package models

import (
	"github.com/rleibl/brauspielhaus/config"
	"reflect"
	"testing"
)

func TestBeersLoad(t *testing.T) {

	c := config.GetConfig()
	b := LoadBeersFromJson(c.JsonPath)

	if len(b) == 0 {
		t.Errorf("No beers read from json")
	}
}

func TestJson(t *testing.T) {

	tmpl := BeerTemplate()
	j := tmpl.ToJson()
	tmpl2 := FromJson([]byte(j))

	// FIXME
	// Fix date handling in general.
	// In this case, the original date has some kind of offset added to
	// it, which makes the comparison fail.
	//     tmpl:  Date: 2020-04-22 09:59:55.39177 +0200 CEST m=+0.001156567
	//     tmpl2: Date: 2020-04-22 09:59:55.39177 +0200 CEST
	tmpl2.Brewdate = tmpl.Brewdate

	if !reflect.DeepEqual(tmpl, tmpl2) {
		tmpl.Print()
		tmpl2.Print()
		t.Errorf("Template -> JSON -> Template Failed")
	}
}
