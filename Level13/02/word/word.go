package word

import "strings"

//UseCount count things
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count count words on a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
