package sprint

func AlphabetMastery(n int) string {
	if n == 0 {
		return ""
	}
	
	if n < 1 || n > 26 {
		return "invalid input"
	}

	var result string
	for i := 0; i < n; i++ {
		result += string('a' + i)
	}

	return result
}