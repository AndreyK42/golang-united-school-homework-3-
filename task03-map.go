package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	length := len(input)
	result = make([]string, 0, length)

	keys := make([]int, 0, length)
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for k := range keys {
		result = append(result, input[k])
	}

	return result
}
