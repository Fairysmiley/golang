package sprint

import "fmt"

func CombN(n int) []string {
	var result []string

	for i := 0; i <= 9; i++ {
		if n > 1 {
			generateCombinations(n-1, fmt.Sprint(i), i+1, &result)
		} else {
			result = append(result, fmt.Sprint(i))
		}
	}

	return result
}

func generateCombinations(n int, current string, start int, result *[]string) {
	if n == 0 {
		*result = append(*result, current)
		return
	}

	for i := start; i <= 9; i++ {
		next := current + fmt.Sprint(i)
		generateCombinations(n-1, next, i+1, result)
	}
}

