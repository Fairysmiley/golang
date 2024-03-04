package sprint

func FromRoman(s string) int {

	var roman = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += roman[rune(s[i])]
		} else if roman[rune(s[i])] < roman[rune(s[i+1])] {
			result -= roman[rune(s[i])]
		} else {
			result += roman[rune(s[i])]
		}

	}

	return result
}
