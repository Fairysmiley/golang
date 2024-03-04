package sprint

func Sqrt(n int) int {
	root := 0

	for i := 0; i*i <= n; i++ {
		if i*i == n {
			root = i
			break
		}
	}

	return root
}
