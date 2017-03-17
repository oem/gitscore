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
	handle(err)

	stats, err := github.GetStats(repos, *token)
	handle(err)

	fmt.Printf("%v\n", stats)
}

func handle(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
