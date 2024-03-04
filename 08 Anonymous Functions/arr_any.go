package sprint

func ArrAny(f func(string) bool, a []string) bool {

	for _, i := range a {
		if f(i) {
			return true
		}
	}
	return false
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
