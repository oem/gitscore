package main

import (
	"flag"
	"fmt"
	"log"

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

	for rank, contributor := range stats {
		fmt.Printf("%3d. %s: %d\n", rank+1, contributor.Name, contributor.Commits)
	}
}
