package problems

import "sort"

func frequencySort(s string) string {
	var res string
	var hash = make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		if v, ok := hash[s[i]]; ok {
			hash[s[i]] = v + 1
		} else {
			hash[s[i]] = 1
		}
	}

	type kv struct {
		key   uint8
		value int
	}
	var p []kv

	for k, v := range hash {
		p = append(p, kv{k, v})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].value > p[j].value
	})
	for _, v := range p {
		for i := 0; i < v.value; i++ {
			res += string(v.key)
		}
	}

	return res
}
