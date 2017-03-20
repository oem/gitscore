package github

import "sort"

type contributor struct {
	Name    string
	Commits int
}

// Contributors is a sortable list of contributor
type Contributors []contributor

func (p Contributors) Len() int           { return len(p) }
func (p Contributors) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Contributors) Less(i, j int) bool { return p[i].Commits < p[j].Commits }

func sortStats(stats map[string]int) Contributors {
	p := make(Contributors, len(stats))
	i := 0
	for k, v := range stats {
		p[i] = contributor{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}
