package sprint

func StrSplitBy(s, sep string) []string {

	if len(s) == 0 {
		return []string{}
	}

	var result []string

	start := 0
	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
		}
	}

	result = append(result, s[start:])

	return result
}
