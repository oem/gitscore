package main

import (
	"flag"
	"log"

	"github.com/oem/gitscore/ansi"
	"github.com/oem/gitscore/github"
)

var token = flag.String("token", "", "github token")
var orga = flag.String("orga", "", "github organisation")
var verbose = flag.Bool("verbose", false, "verbose logging")

func init() {
	flag.BoolVar(verbose, "v", false, "verbose logging")
}

func main() {
	flag.Parse()

	repos, err := github.Repos(*orga, *token)
	if err != nil {
		log.Fatal(err)
	}
	stats := github.GetStats(*orga, repos, *token, *verbose)

	err = ansi.Draw(stats)
	if err != nil {
		log.Fatal(err)
	}
}
