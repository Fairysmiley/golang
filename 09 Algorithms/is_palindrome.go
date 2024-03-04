package sprint

func IsPalindrome(s string) bool {

	newString := cleanstring(s)
	reversedstring := StrReverse(newString)
	return newString == reversedstring
}

func TolowerCase(r rune) rune {

	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func StrReverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func cleanstring(s string) string {

	var newstring []rune
	for _, i := range s {
		if (i > '0' || i < '9') && (i > 'A' || i < 'Z') && (i > 'a' || i < 'z') {
			newstring = append(newstring, TolowerCase(i))
		}
	}
	return string(newstring)

}
