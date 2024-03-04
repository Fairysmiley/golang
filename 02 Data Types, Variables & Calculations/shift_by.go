package sprint

func ShiftBy(r rune, step int) rune {
	if r < 'a' || r > 'z' {
		return r
	}

	shifted := ((int(r) - 'a' + step) % 26) + 'a'

	if shifted < 'a' {
		shifted += 26
	}

	return rune(shifted)
}
