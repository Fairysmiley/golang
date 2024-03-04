package sprint

func ToThePowerRecursive(n, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	result := ToThePowerRecursive(n, power-1)

	if n == 0 || (power > 0 && n > 0 && result > (1<<31-1)/n) {
		return 0
	}

	if n < 0 && (power%2 != 0 && -result < (-1<<31)/(-n) || result > (1<<31-1)/(-n)) {
		return 0
	}

	return result * n
}
