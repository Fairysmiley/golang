package sprint

func SplitWhitespaces(s string) []string {

	var allWords []string
	var currentWord string

	for _, i := range s {
		if i == 9 || i == 10 || i == 11 || i == 32 {
			if currentWord != "" {
				allWords = append(allWords, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(i)
		}
	}

	if currentWord != "" {
		allWords = append(allWords, currentWord)
	}

	return allWords
}

