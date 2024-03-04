package sprint

func StrLength(s string) []int {

	runeCount := 0
	for range s {
		runeCount += 1
	}
	byteLength := len(s)

	result := []int{runeCount, byteLength}
	return result
}
