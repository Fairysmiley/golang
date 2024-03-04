package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	var result [][]int

	for _, subArr := range arr {
		sum := 0
		for _, value := range subArr {
			sum += value
		}
		if sum >= limit {
			result = append(result, subArr)
		}
	}
	if len(result) == 0 {
		return [][]int{}
	}
	return result
}
