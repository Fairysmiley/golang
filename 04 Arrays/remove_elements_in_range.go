package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	length := len(arr)

	if from > to {
		from, to = to, from
	}

	if from <= 0 && to > length {
		return []float64{}
	}

	if to < 0 || from >= length {
		return arr
	}

	if from < 0 && to < 0 {

		return arr
	}

	if from > length && to > length {

		return arr
	}

	if from < 0 && to < length {

		return arr[to:]
	}

	if to > length && from >= 0 {
		return arr[:from]
	}

	if from < 0 {
		from = -1
	}

	if to >= length {
		to = length - 1
	}

	result := append(arr[:from], arr[to:]...)
	return result
}
