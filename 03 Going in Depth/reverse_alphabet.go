package sprint

func ReverseAlphabet(step int) string {

	if step <= 0 {
		step = 1
	}
	var result string
	for i := 'z'; i >= 'a'; i -= rune(step) {
		result += string(i)
	}
	return result
}
