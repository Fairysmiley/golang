package sprint

func SimpleStrToInt(s string) int {
	result := 0

	for _, digit := range s {

		if digit >= '0' && digit <= '9' {
			num := int(digit - '0')
			result = result*10 + num
		} else {

			return 0
		}
	}

	return result

}

