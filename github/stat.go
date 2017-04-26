package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type author struct {
	Total  int `json:"total"`
	Author struct {
		Name string `json:"login"`
	} `json:"author"`
}

// GetStats is getting the list of contributors for all the repos of the organisation in parallel
// returns a sorted list of contributors <github.Contributors>
// Errors will only be logged but otherwise ignored
func GetStats(orga string, repos []string, token string) Contributors {
	c := make(chan []author, len(repos))

	var authors []author

	for _, repo := range repos {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/stats/contributors", orga, repo)
		go func() {
			for {
				proj, err := getStat(token, url)
				if err == nil {
					c <- proj
					break
				}
				log.Println(err)
				log.Printf("requeuing %s\n", url)
			}
		}()
	}

	for i := 0; i < len(repos); i++ {
		proj := <-c
		authors = append(authors, proj...)
	}

	stats := sumStats(authors)
	return sortStats(stats)
}

func getStat(token string, url string) ([]author, error) {
	var authors []author

	timeout := time.Duration(8 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("oem", token)
	resp, err := client.Do(req)
	if err != nil {
		return authors, err
	}

	if resp.StatusCode == 204 {
		return authors, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return authors, err
	}

	err = json.Unmarshal(body, &authors)
	return authors, err
}

func sumStats(authors []author) map[string]int {
	stats := map[string]int{}
	for _, author := range authors {
		stats[author.Author.Name] += author.Total
	}

	return stats
}
