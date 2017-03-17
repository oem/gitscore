package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oem/highscore/github"
)

var token = flag.String("token", "", "github token")

func main() {
	flag.Parse()

	repos, err := github.GetRepos(*token)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	stats, err := github.GetStats(repos, *token)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", stats)
}
