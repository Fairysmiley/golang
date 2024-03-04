package sprint

func ReverseAlphabetValue(ch rune) rune {
	if 'a' > ch || 'z' < ch {
		return ch
	}
	var reverse rune
	reverse = 'a' + ('z' - ch)
	return rune(reverse)
}
