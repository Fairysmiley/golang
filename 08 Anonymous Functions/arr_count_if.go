package sprint

func ArrCountIf(f func(string) bool, tab []string) int {

	countTrue := 0

	for _, i := range tab {
		if f(i) {
			countTrue++
		}
	}
	return countTrue
}

func IsLower(s string) bool {

	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {

	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool {

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {

	for i := 0; i < len(s); i++ {
		if (s[i] < '0' || s[i] > '9') && (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') {
			return false
		}
	}
	return true

}
