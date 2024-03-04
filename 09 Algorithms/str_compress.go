package sprint

import "fmt"

func StrCompress(input string) string {
	if len(input) <= 1 {
		return input
	}

	count := 1
	result := ""

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count != 1 {
				result += fmt.Sprintf("%d", count)
			}
			result += string(input[i-1])
			count = 1
		}
	}
	if count > 1 {
		result += fmt.Sprintf("%d", count)
	}

	return result + string(input[len(input)-1])

}
