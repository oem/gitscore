package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/oem/gitscore/github"
)

func main() {
	var (
		token   = flag.String("token", "", "github token")
		orga    = flag.String("orga", "", "github organisation")
		verbose = flag.Bool("verbose", false, "verbose logging")
	)
	flag.BoolVar(verbose, "v", false, "verbose logging")
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
