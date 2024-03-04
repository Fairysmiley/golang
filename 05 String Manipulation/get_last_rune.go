package sprint

func GetLastRune(s string) rune {

	var length int
	//length := len(s)

	for range s {
		length++
	}

	if length > 0 {
		return []rune(s)[length-1]
	}
	return 0
}

