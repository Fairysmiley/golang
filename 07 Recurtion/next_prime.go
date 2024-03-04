package sprint

func NextPrime(n int) int {

	for !IsPrime(n) {
		n++
	}
	return n
}

func IsPrime(n int) bool {

	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
