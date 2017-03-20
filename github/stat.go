package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type contributor struct {
	Total  int `json:"total"`
	Author struct {
		Name string `json:"login"`
	} `json:"author"`
}

type result struct {
	contributors []contributor
	err          error
}

// GetStats is getting the list of contributors for all the organisations repos in parallel
// returns map[String]int, key is the contributors name, value the sum of commits
// Errors will only be logged but otherwise ignored
func GetStats(orga string, repos []string, token string) pairlist {
	c := make(chan result, len(repos))

	var contributors []contributor

	for _, repo := range repos {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/stats/contributors", orga, repo)
		go func() {
			proj, err := getStat(token, url)
			c <- result{contributors: proj, err: err}
		}()
	}

	for i := 0; i < len(repos); i++ {
		result := <-c
		if result.err != nil {
			log.Printf("%v\n", result.err)
			continue
		}
		contributors = append(contributors, result.contributors...)
	}

	stats := sumStats(contributors)
	return sortStats(stats)
}

func getStat(token string, url string) ([]contributor, error) {
	var contributors []contributor

	timeout := time.Duration(8 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("oem", token)
	resp, err := client.Do(req)
	if err != nil {
		return contributors, err
	}

	if resp.StatusCode == 204 {
		return contributors, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return contributors, err
	}

	err = json.Unmarshal(body, &contributors)
	log.Printf("%s: %v", url, contributors)
	return contributors, err
}

func sumStats(contributors []contributor) map[string]int {
	stats := map[string]int{}
	for _, contributor := range contributors {
		stats[contributor.Author.Name] += contributor.Total
	}

	return stats
}
