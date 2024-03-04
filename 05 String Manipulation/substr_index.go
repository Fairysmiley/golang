package sprint

func SubstrIndex(s string, toFind string) int {

	if len(toFind) == 0 {
		return 0
	}

	if len(s) == 0 {
		return -1
	}

	for i := 0; i < len(s)-len(toFind)+1; i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}

	return -1
}
