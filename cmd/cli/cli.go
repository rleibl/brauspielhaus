package main

import (
	"fmt"
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
`
	fmt.Print(help)

	os.Exit(status)
}

func printTemplate(s []string) {
	models.PrintBeerExample()
}
