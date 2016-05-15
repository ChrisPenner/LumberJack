package main

// Filter over strings
func Filter(items []string, f func(string) bool) []string {
	var result []string
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
