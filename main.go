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
	var contributors []github.Contributor

	repos, err := github.GetRepos(*token)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, repo := range repos {
		url := fmt.Sprintf("https://api.github.com/repos/njiuko/%s/stats/contributors", repo)
		fmt.Println(repo)
		projContributors, err := github.GetStat(*token, url)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		contributors = append(contributors, projContributors...)
	}
	stats, err := github.SumStats(contributors)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", stats)
}
