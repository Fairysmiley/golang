package sprint

func AlphaNumber(n int) string {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	numStr := ""

	if n == 0 {
		numStr = "0"
	}

	for n > 0 {
		digit := '0' + byte(n%10)
		numStr = string(digit) + numStr
		n /= 10
	}

	result := ""
	for _, digit := range numStr {
		digitValue := int(digit - '0')
		result += string('a' + digitValue)
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
