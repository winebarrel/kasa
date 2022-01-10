package utils

import "sort"

func Uniq(ary []string) []string {
	set := make(map[string]struct{})

	for _, k := range ary {
		set[k] = struct{}{}
	}

	newAry := make([]string, len(set))
	i := 0

	for k := range set {
		newAry[i] = k
		i++
	}

	sort.Slice(newAry, func(i, j int) bool { return newAry[i] < newAry[j] })

	return newAry
}
