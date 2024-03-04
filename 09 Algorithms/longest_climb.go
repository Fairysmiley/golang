package sprint

func LongestClimb(arr []int) []int {

	result := []int{}
	max := 1
	length := 1
	maxIndex := 0

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] {
			length++
		} else {
			if max < length {
				max = length
				maxIndex = i - max
			}
			length = 1
		}
	}
	if max < length {
		max = length
		maxIndex = len(arr) - max
	}
	result = append(result, arr[maxIndex:maxIndex+max]...)
	return result
}
