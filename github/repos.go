package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type repo struct {
	Name string `json:"name"`
}

// GetRepos gets all the repos for an organisation via the github api
func GetRepos(orga string, token string) ([]string, error) {
	var names []string
	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos?type=sources", orga)

	// TODO this will not scale ofc, once we have more than 90 repos (30 repos per page)
	for i := 1; i <= 3; i++ {
		repos, err := paginatedRepos(token, fmt.Sprintf("%s&page=%d", url, i))
		if err != nil {
			return names, err
		}
		names = append(names, repos...)
	}
	return names, nil
}

func paginatedRepos(token string, url string) ([]string, error) {
	var repoNames []string
	var err error

	timeout := time.Duration(4 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("oem", token)
	resp, err := client.Do(req)
	if err != nil {
		return repoNames, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return repoNames, err
	}

	return extractNames(body)
}

func extractNames(body []byte) ([]string, error) {
	var names []string
	var repos []repo

	err := json.Unmarshal(body, &repos)
	if err != nil {
		return names, err
	}

	for _, repo := range repos {
		names = append(names, repo.Name)
	}
	return names, err
}
