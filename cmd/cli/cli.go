package main

import (
	"fmt"
	"github.com/rleibl/brauspielhaus/config"
	"github.com/rleibl/brauspielhaus/models"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		helpAndExit(-1)
	}

	cmd := os.Args[1]

	opts := os.Args[2:]

	switch cmd {
	case "help":
		helpAndExit(0)
	case "template":
		printTemplate(opts)
	case "fill":
		fillBeerJson(opts)
	case "list":
		listBeers(opts)
	case "print":
		printBeers(opts)
	default:
		fmt.Printf("unknown command: %s\n", cmd)
		helpAndExit(-1)
	}
}

func helpAndExit(status int) {

	help := `
usage: cli <command> [options]")

available commands

    help
        print this help and exit

    template
        print a default beer template

    list
        print a summary of all available beers
    
    print
        print all beers

    fill <beer.json>
        Update all automatically calculated fields in <beer.json> .
	Print the resulting json to stdout
`
	fmt.Print(help)

	os.Exit(status)
}

func printTemplate(s []string) {
	models.PrintBeerTemplate()
}

func listBeers(s []string) {

	// Use load from json directly. This way, we can give the path
	// directly.
	// Alternatively:
	//     beers := db.getBeers()
	c := config.GetConfig()
	beers := models.LoadBeersFromJson(c.JsonPath)

	for i, b := range beers {
		fmt.Printf("%d: %s, %s\n", i+1, b.Name, b.Brewdate)
	}
}

func printBeers(s []string) {

	// Use load from json directly. This way, we can give the path
	// directly.
	// Alternatively:
	//     beers := db.getBeers()
	c := config.GetConfig()
	beers := models.LoadBeersFromJson(c.JsonPath)

	for _, b := range beers {
		b.Print()
	}
}

func fillBeerJson(opts []string) {
	if len(opts) < 1 {
		fmt.Println("No filename given for 'fill' command")
		helpAndExit(-1)
	}

	filename := opts[0]
	b := models.LoadBeerFromJson(filename)
	fmt.Println(b.ToJson())
}
