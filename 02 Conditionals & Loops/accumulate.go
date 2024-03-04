package sprint

func Accumulate(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	if n < 0 {
		return 0
	}
	return sum
}

