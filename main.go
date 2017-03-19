package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oem/highscore/github"
)

var token = flag.String("token", "", "github token")
var orga = flag.String("orga", "", "github organisation")

func main() {
	flag.Parse()

	repos, err := github.GetRepos(*orga, *token)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	stats := github.GetStats(*orga, repos, *token)

	fmt.Printf("%v\n", stats)
}
