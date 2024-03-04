package sprint

func StrToInt(s string) int {
	result, sign := 0, 1
	started := false

	for _, char := range s {
		if char == ' ' {
			if started {
				return 0
			}
			continue
		}

		if char == '-' || char == '+' {
			if started {
				break
			}
			if char == '-' {
				sign = -1
			}
			started = true
		} else if char >= '0' && char <= '9' {
			num := int(char - '0')
			result = result*10 + num
			started = true
		} else {
			result = 0
			break
		}
		if started && char == ' ' {
			result = 0
			break
		}
	}

	return result * sign
}
