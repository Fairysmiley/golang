package sprint

func ToCapitalCase(s string) string {
	var result string
	var capitalizeNext bool
	firstChar := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			if capitalizeNext || firstChar {
				if ch >= 'a' && ch <= 'z' {
					ch = ch - 'a' + 'A'
				}
				capitalizeNext = false
				firstChar = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					ch = ch - 'A' + 'a'
				}
			}
			result += string(ch)
		} else {
			capitalizeNext = true
			result += string(ch)
		}
	}

	return result
}
