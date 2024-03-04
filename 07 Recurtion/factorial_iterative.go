package sprint

func FactorialIterative(n int) int {

	if n < 0 {
		return 0
	}

	result := 1

	for i := 1; i <= n; i++ {
		if result > (1<<31-1)/i { //1 << 31 is a bitwise left shift of 1 by 31 positions. 1<<31 - 1 is equal to 2^31 - 1, which is the maximum positive value representable by a signed 32-bit integer. In Go, this value is often used as a constant named math.MaxInt32 in the math package.
			//If result multiplied by the loop variable i exceeds this value, it indicates an overflow, and the function returns 0.
			return 0
		}
		result *= i
	}

	return result
}
