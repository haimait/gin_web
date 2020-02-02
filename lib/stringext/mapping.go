package stringext

import (
	"sort"
)

// 给key为string类型的map做key排序
func MapKeySort(mapping map[string][]string) []string {
	keys := make([]string, len(mapping))
	i := 0
	for k := range mapping {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
