package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type contributor struct {
	Total  int `json:"total"`
	Author struct {
		Name string `json:"login"`
	} `json:"author"`
}

func GetStats(repos []string, token string) (map[string]int, error) {
	var contributors []contributor
	for _, repo := range repos {
		url := fmt.Sprintf("https://api.github.com/repos/njiuko/%s/stats/contributors", repo)
		projContributors, err := getStat(token, url)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		contributors = append(contributors, projContributors...)
	}

	return sumStats(contributors)
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
	if err != nil {
		return contributors, err
	}

	return contributors, err
}

func sumStats(contributors []contributor) (map[string]int, error) {
	stats := map[string]int{}
	for _, contributor := range contributors {
		stats[contributor.Author.Name] += contributor.Total
	}

	return stats, nil
}
