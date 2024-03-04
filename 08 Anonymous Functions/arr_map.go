package sprint

func ArrMap(f func(int) bool, a []int) []bool {

	var result []bool

	for _, i := range a {
		result = append(result, f(i))
	}

	return result
}

func IsNegative(n int) bool {
	return n < 0
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
