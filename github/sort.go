package github

import "sort"

type pair struct {
	key   string
	value int
}

type pairlist []pair

func (p pairlist) Len() int           { return len(p) }
func (p pairlist) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairlist) Less(i, j int) bool { return p[i].value < p[j].value }

func sortStats(stats map[string]int) pairlist {
	p := make(pairlist, len(stats))
	i := 0
	for k, v := range stats {
		p[i] = pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}
