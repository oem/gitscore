package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type client interface {
	getPage(string) ([]byte, string, error)
}

type apiClient struct{}

type repo struct {
	Name string `json:"name"`
}

// Repos is a thin wrapper around getRepos, using the ApiClient
// returns []string of repos and error
func Repos(orga string, token string) ([]string, error) {
	client := apiClient{}
	return getRepos(client, orga, token)
}

func getRepos(client client, orga string, token string) ([]string, error) {
	var names []string
	var repos []byte
	var err error
	next := fmt.Sprintf("https://api.github.com/orgs/%s/repos?type=sources&access_token=%s", orga, token)

	for {
		repos, next, err = client.getPage(next)
		if err != nil {
			return nil, err
		}

		repoNames, err := extractNames(repos)
		if err != nil {
			return nil, err
		}
		names = append(names, repoNames...)

		if len(next) == 0 {
			break
		}
	}
	return names, nil
}

func (api apiClient) getPage(url string) ([]byte, string, error) {
	var err error
	var next string

	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	link := resp.Header.Get("Link")
	next, err = extractNext(link)
	if err != nil {
		return nil, "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, next, err
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

func extractNext(link string) (string, error) {
	rp, err := regexp.Compile("<(http.+?)>; rel=\"next\"")
	if err != nil {
		return "", err
	}
	nextLink := rp.FindSubmatch([]byte(link))
	if len(nextLink) <= 0 {
		return "", nil
	}
	return string(nextLink[1]), nil
}
