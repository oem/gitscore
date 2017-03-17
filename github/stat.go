package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Contributor struct {
	Total  int `json:"total"`
	Author struct {
		Name string `json:"login"`
	} `json:"author"`
}

func GetStat(token string, url string) ([]Contributor, error) {
	var contributors []Contributor

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

func SumStats(contributors []Contributor) (map[string]int, error) {
	stats := map[string]int{}
	for _, contributor := range contributors {
		stats[contributor.Author.Name] += contributor.Total
	}

	return stats, nil
}
