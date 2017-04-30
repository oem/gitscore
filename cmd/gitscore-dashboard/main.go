package main

import (
	"flag"
	"log"

	"github.com/oem/gitscore/ansi"
	"github.com/oem/gitscore/github"
)

var token = flag.String("token", "", "github token")
var orga = flag.String("orga", "", "github organisation")

func main() {
	flag.Parse()

	repos, err := github.Repos(*orga, *token)
	if err != nil {
		log.Fatal(err)
	}
	stats := github.GetStats(*orga, repos, *token)

	err = ansi.Draw(stats)
	if err != nil {
		log.Fatal(err)
	}
}
