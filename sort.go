package balance

import (
	"sort"
)

func mapSortByValue(m map[string]int) (result []string) {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	result = make([]string, len(ss))
	for i, v := range ss {
		result[i] = v.Key
	}
	return
}
