package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {

	n := len(arr)
	ascending := true
	for i := 0; i < n-1; i++ {
		if f(arr[i], arr[i+1]) > 0 {
			ascending = false
			break
		}
	}

	descending := true
	for i := 0; i < n-1; i++ {
		if f(arr[i], arr[i+1]) < 0 {
			descending = false
			break
		}
	}

	return ascending || descending
}

func StrCompare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}
