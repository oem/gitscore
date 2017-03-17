package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type contributor struct {
	Total  int `json:"total"`
	Author struct {
		Name string `json:"login"`
	} `json:"author"`
}

func GetStats(repos []string, token string) (pairlist, error) {
	var contributors []contributor
	for _, repo := range repos {
		url := fmt.Sprintf("https://api.github.com/repos/njiuko/%s/stats/contributors", repo)
		projContributors, err := getStat(token, url)
		if err != nil {
			return nil, err
		}
		contributors = append(contributors, projContributors...)
	}

	stats, err := sumStats(contributors)
	if err != nil {
		return nil, err
	}
	return sortStats(stats), nil
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
	return contributors, err
}

func sumStats(contributors []contributor) (map[string]int, error) {
	stats := map[string]int{}
	for _, contributor := range contributors {
		stats[contributor.Author.Name] += contributor.Total
	}

	return stats, nil
}
